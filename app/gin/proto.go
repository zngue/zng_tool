package gin

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var ProtoCommand = &cobra.Command{
	Use:   "gb",
	Short: "Generate gin code through PB zng gb test/abc.proto ",
	Long:  `通过pb 生成 gin 的代码`,
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("参数错误")
		return
	}
	protoPath := args[0]
	if !strings.Contains(protoPath, ".proto") {
		fmt.Println("请输入正确的文件")
		return
	}
	Exec(protoPath)
}

func Exec(proto string) {
	//proto := "api/gin-pb/v1/gin-pb.proto"
	path, err := exec.LookPath("protoc-gen-gin-http")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if path == "" {
		fmt.Println("请安装protoc-gen-gin-http 插件")
		return
	}
	fmt.Println("文件所在路径是：", path)
	input := []string{
		"--proto_path=.",
	}
	//判断第三方依赖
	if pathExists("third_party") {
		input = append(input,
			"--proto_path=./third_party",
		)
	}
	inputExt := []string{
		"--go_out=paths=source_relative:.",
		"--gin-http_out=paths=source_relative:.",
	}
	input = append(input, inputExt...)
	var protoBytes []byte
	protoBytes, err = os.ReadFile(proto)
	if err == nil && len(protoBytes) > 0 {
		if ok, _ := regexp.Match(`\n[^/]*(import)\s+"validate/validate.proto"`, protoBytes); ok {
			input = append(input, "--validate_out=lang=go,paths=source_relative:.")
		}
	}
	input = append(input, proto)
	fd := exec.Command("protoc", input...)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = "."
	if err = fd.Run(); err != nil {
		fmt.Println(err)
		return
	}
	//判断是否存在validate 目录存在删除
	if pathExists("validate") {
		err = os.RemoveAll("validate")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	//google
	if pathExists("google") {
		err = os.RemoveAll("google")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Printf("proto: %s\n", proto)
	fmt.Println("over")
}
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
