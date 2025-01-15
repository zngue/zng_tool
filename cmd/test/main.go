package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	proto := "api/test/v1/test.proto"
	path, err := exec.LookPath("protoc-gen-gin-http")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	input := []string{
		"--proto_path=.",
	}
	fmt.Println("文件所在路径是：", path)

	input = append(input,
		"--proto_path=./third",
	)
	inputExt := []string{
		"--go_out=paths=source_relative:.",
		"--gin-http_out=paths=source_relative:.",
	}
	input = append(input, inputExt...)

	//protoBytes, err := os.ReadFile(proto)
	//if err == nil && len(protoBytes) > 0 {
	//	if ok, _ := regexp.Match(`\n[^/]*(import)\s+"validate/validate.proto"`, protoBytes); ok {
	//		input = append(input, "--validate_out=lang=go,paths=source_relative:.")
	//	}
	//}
	input = append(input, proto)
	fd := exec.Command("protoc", input...)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = "."
	if err = fd.Run(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("proto: %s\n", proto)
	fmt.Println(path, err)
}
