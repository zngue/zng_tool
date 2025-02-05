package project

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zngue/zng_tool/app"
	"io"
	"os"
	"os/exec"
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
			fmt.Println(fmt.Sprintf("cd %s", name))
			fmt.Println("go mod tidy")
			//判断是否是windows
			if IsWindows() {
				command := exec.Command("cmd", "/c", fmt.Sprintf("cd %s  && go mod tidy", name))
				var stdout io.ReadCloser
				stdout, err = command.StdoutPipe() // 标准输出
				if err != nil {
					fmt.Println("222", err)
				}
				if err != nil {
					fmt.Println(err)
					return
				}
				err = command.Start() // 执行命令
				if err != nil {
					fmt.Println("333", err)
				}
				inputReader := bufio.NewReader(stdout)
				for {
					var line, errs = inputReader.ReadString('\n') // 一行一行地读取数据
					if errs == io.EOF {                           //读取完成
						break
					}
					fmt.Println(line)
				}
				if err = command.Wait(); err != nil {
					fmt.Println("8888", err)
				}

				//标准输出
			} else {
				err := exec.Command("bash", "-c", fmt.Sprintf("cd %s  && go mod tidy", name)).Run()
				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}
)

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
