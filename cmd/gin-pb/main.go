package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {

	proto := "api/test/v1/test.proto"
	path, err := exec.LookPath("protoc-gen-gin-tmp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if path == "" {
		fmt.Println("请安装protoc-gen-gin-tmp 插件")
		return
	}
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
		"--gin-tmp_out=paths=source_relative:.",
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
