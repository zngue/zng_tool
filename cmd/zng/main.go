package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zngue/zng_tool/app/config"
	"github.com/zngue/zng_tool/app/project"
	"github.com/zngue/zng_tool/app/proto"
)

var rootCmd = &cobra.Command{
	Use:     "zng",
	Short:   "zng: An elegant toolkit for Go microservices.",
	Long:    `zng: An elegant toolkit for Go microservices.`,
	Version: "1.0.0",
}

func main() {
	var err error
	config.Run()
	rootCmd.AddCommand(
		project.CommandProject,
		config.Init(),
		proto.Init(),
	)
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
	return
}
