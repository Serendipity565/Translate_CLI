package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionFlag bool
var version = "v1.0"

var rootCmd = &cobra.Command{
	Use:   "translate",
	Short: "A simple and efficient translation tool",
	Long: `translate is a powerful yet easy-to-use translation tool. 
It supports multiple languages and provides fast, accurate translations directly from the command line.
Complete documentation is available at `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("Current version:", version)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo...")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number and exit")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
