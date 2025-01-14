package proto

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
)

var command = &cobra.Command{
	Use:   "p",
	Long:  `zng: zng proto`,
	Short: "proto pb生成",
}

func Init() *cobra.Command {
	command.AddCommand(
		Client(),
	)
	return command
}
func Client() *cobra.Command {
	return &cobra.Command{
		Use:   "c",
		Short: "client",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("参数错误")
				return
			}
			var (
				httpKey     = ""
				grpcKey     = "protoc --proto_path=./third  --proto_path=.  --go_out=. --go-grpc_out=. "
				validateKey = ""
			)
			var isWindows = IsWindows()
			goBin := os.Getenv("GOBIN")
			if goBin == "" {
				fmt.Println("请设置GOBIN go env -w GOBIN=G:\\go\\bin")
				return
			}
			fmt.Println("goBin:", goBin)
			var binGinHttp string
			if isWindows {
				binGinHttp = fmt.Sprintf("%s\\protoc-gen-protoc-gen-gin-http.exe", goBin)
			} else {
				binGinHttp = fmt.Sprintf("%s/protoc-gen-protoc-gen-gin-http", goBin)
			}
			//获取go bin 所在目录
			if isWindows {
				httpKey = "protoc --proto_path=./third  --proto_path=. --plugin=protoc-gen-custom=" + binGinHttp + " --custom_out=.  "
				validateKey = "protoc --proto_path=./third  --proto_path=. --plugin=protoc-gen-validate.exe  --go_out=. --validate_out=lang=go:. "
			} else {
				httpKey = "protoc --proto_path=./third  --proto_path=. --plugin=protoc-gen-custom=" + binGinHttp + " --custom_out=. "
				validateKey = "protoc --proto_path=./third  --proto_path=. --plugin=protoc-gen-validate  --go_out=. --validate_out=lang=go:. "
			}
			Run(httpKey, args[0])
			Run(grpcKey, args[0])
			Run(validateKey, args[0])
			//判断文件夹是否存在
			var path = "google.golang.org"
			deleteFolderIfExists(path)

		},
	}
}
func deleteFolderIfExists(path string) {
	// 判断文件夹是否存在
	info, err := os.Stat(path)
	if err != nil {
		// 文件夹不存在，直接返回
		if os.IsNotExist(err) {
			return
		}
		// 如果是其他错误，打印并返回
		fmt.Println("无法获取文件夹信息:", err)
		return
	}

	// 判断是否是目录
	if info.IsDir() {
		// 删除文件夹及其内容
		err = os.RemoveAll(path)
		if err != nil {
			fmt.Println("删除文件夹失败:", err)
			return
		}
		fmt.Println("文件夹已删除:", path)
	} else {
		fmt.Println("路径不是文件夹:", path)
	}
}

// IsWindows 判断是windows还是linux
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
func Run(httpKey string, file string) {
	if IsWindows() {
		runKey := fmt.Sprintf("%s %s", httpKey, file)
		fmt.Println(runKey)
		httpKeyCommand := exec.Command("cmd", "/c", runKey)
		output, err := httpKeyCommand.Output()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(output))
		}
	} else {
		httpKeyCommand := exec.Command(fmt.Sprintf("%s %s", httpKey, file))
		output, err := httpKeyCommand.Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(output))
		}
	}
}
