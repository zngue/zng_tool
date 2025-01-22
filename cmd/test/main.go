package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {

	//接受第一个参数
	if len(os.Args) <= 1 {
		fmt.Println("请输入参数1")
		return
	}
	var proto = os.Args[1]
	if proto == "" {
		fmt.Println("请输入proto文件名2")
		return
	}
	//proto := "api/test/v1/test.proto"
	path, err := exec.LookPath("protoc-gen-gin-http")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件所在路径是：", path)
	path, err = exec.LookPath("protoc-gen-kratos-temp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件所在路径是：", path)
	input := []string{
		"--proto_path=.",
	}

	input = append(input,
		"--proto_path=./third_party",
	)
	inputExt := []string{
		"--go_out=paths=source_relative:.",
		"--gin-http_out=paths=source_relative:.",
	}
	input = append(input, inputExt...)
	protoBytes, err := os.ReadFile(proto)
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
		os.RemoveAll("validate")
	}
	//google
	if pathExists("google") {
		os.RemoveAll("google")
	}
	fmt.Printf("proto: %s\n", proto)
	fmt.Println(path, err)
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
