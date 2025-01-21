package template

import "github.com/spf13/cobra"

var command = &cobra.Command{
	Use:   "tmp",
	Long:  `zng: zng proto`,
	Short: "template zng tmp ",
}

func Init() *cobra.Command {
	command.AddCommand(kratosTmp)
	return command
}
