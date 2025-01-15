package build

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CommandBuild = &cobra.Command{
	Use:  "b",
	Long: `zng: zng build`,
}
var CommandUrl = &cobra.Command{
	Use:  "url",
	Long: `zng: zng build url`,
	Run: func(cmd *cobra.Command, args []string) {
		//判断Url
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}
		//判断Url
		url := args[0]
		fmt.Println(url)

	},
}
var CommandName = &cobra.Command{
	Use:  "name",
	Long: `zng: zng build name`,
	Run: func(cmd *cobra.Command, args []string) {
		//判断Url
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				return
			}
			return
		}

	},
}

func Set(url string) {

}
