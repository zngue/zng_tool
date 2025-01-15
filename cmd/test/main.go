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
		fmt.Println(err)
		return
	}
	input := []string{
		"--proto_path=.",
	}
	inputExt := []string{
		"--proto_path=./third",
		"--go_out=paths=source_relative:.",
		"--go-grpc_out=paths=source_relative:.",
		"--gin-http_out=paths=source_relative:.",
	}
	input = append(input, inputExt...)
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
