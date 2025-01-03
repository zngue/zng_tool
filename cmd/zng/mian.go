package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zngue/zng_tool/app"
	"os"
)

func main() {
	//name   = "user-service"
	var rootCmd = &cobra.Command{
		Use:     "zng",
		Short:   "zng: An elegant toolkit for Go microservices.",
		Long:    `zng: An elegant toolkit for Go microservices.`,
		Version: "1.0.0",
	}
	rootCmd.AddCommand(
		&cobra.Command{
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
			},
		},
	)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
	return
}
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
	return
}
