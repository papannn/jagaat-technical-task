package main

import (
	"github.com/spf13/cobra"
	"jagaat-technical-task/cmd/get_user_data_command"
	"jagaat-technical-task/config"
)

var rootCmd = &cobra.Command{
	Long: `
	Welcome to jagaat-technical-task cli
`,
}

func main() {
	rootCmd.AddCommand(get_user_data_command.Command)
	rootCmd.Execute()
}

func init() {
	config.InitializeConfig()
}
