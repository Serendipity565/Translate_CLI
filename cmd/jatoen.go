package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var jatoenCmd = &cobra.Command{
	Use:   "ja-en <text>", // 子命令名称
	Short: "Translate Japanese to English",
	Args:  cobra.ExactArgs(1), // 确保必须传入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0] // 从命令行参数获取输入

		translatesever, err := NewServiceContainer()
		res, err := translatesever.Translator.Translate(text, "ja", "en")
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	// 注册子命令到根命令
	rootCmd.AddCommand(jatoenCmd)
}
