package project

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zngue/zng_tool/app"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	CommandProject = &cobra.Command{
		Use:  "new",
		Long: "创建一个项目",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("请输入项目名称")
			}
			fmt.Println(fmt.Sprintf("开始创建项目：%s", args[0]))
			var (
				name   = args[0]
				branch = "master"
			)
			if len(args) >= 2 {
				branch = args[1]
			}
			var err = run(name, branch)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("项目创建成功")
			if IsWindows() {
				fmt.Println(fmt.Sprintf("cd %s", name))
				fmt.Println("go mod tidy")
				ChangeDir(name)
				execCmd := exec.Command("go", "mod", "tidy")
				execCmd.Stdout = os.Stdout
				execCmd.Stderr = os.Stderr
				if err := execCmd.Run(); err != nil {
					fmt.Println(err)
					return
				}
			} else {
				execCmd := exec.Command("sh", fmt.Sprintf("cd %s && go mod tidy", name))
				execCmd.Stdout = os.Stdout
				execCmd.Stderr = os.Stderr
				if err := execCmd.Run(); err != nil {
					fmt.Println(err)
					return
				}
			}

		},
	}
)

func ChangeDir(name string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}
	// 目标目录（相对于当前目录）
	targetDir := filepath.Join(currentDir, name)
	// 切换目录
	err = os.Chdir(targetDir)
	if err != nil {
		fmt.Println("切换目录失败:", err)
		return
	}
	newDir, _ := os.Getwd()
	fmt.Println("当前目录:", newDir)
}

func DoCmd(md string) {
	execCmd := exec.Command("cmd", "-c", md)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

//執行shell指令 区分windows 和Linux

func run(name string, branch string) (err error) {
	var (
		url = "https://gitee.com/zngue_mic/zng_layout.git"
	)
	//判断文件夹是否存在
	_, err = os.Stat(name)
	if err == nil {
		fmt.Println(fmt.Sprintf("改项目已存在：%s", name))
		return
	}
	repo := app.NewRepo(url, branch)
	ctx := context.Background()
	err = repo.RemoTo(
		ctx,
		name,
		[]string{".git", ".github", "go.sum"},
		[]string{
			"github.com/zngue/zng_layout",
			fmt.Sprintf("github.com/zngue_mic/%s", name),
			"zng_layout",
			name,
		},
	)
	//cd name git pull go mod tidy

	return
}
