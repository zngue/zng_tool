package service

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestAcc(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}
	// 目标目录（相对于当前目录）
	targetDir := filepath.Join(currentDir, "api")
	// 切换目录
	err = os.Chdir(targetDir)
	if err != nil {
		fmt.Println("切换目录失败:", err)
		return
	}
	newDir, _ := os.Getwd()
	fmt.Println("当前目录:", newDir)

}
