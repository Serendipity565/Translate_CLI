package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionFlag bool
var version = "v1.0"

var rootCmd = &cobra.Command{
	Use:   "tran",
	Short: "A simple and efficient translation tool",
	Long: `translate is a powerful yet easy-to-use translation tool. 
It supports multiple languages and provides fast, accurate translations directly from the command line.
Complete documentation is available at https://github.com/Serendipity565/Translate_CLI`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("Current version:", version)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // 显示帮助信息
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number and exit")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
