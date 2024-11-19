package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var entozhCmd = &cobra.Command{
	Use:   "en-zh", // 子命令名称
	Short: "Translate English to Chinese",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取标志参数
		text, _ := cmd.Flags().GetString("message")
		if text == "" {
			fmt.Println("No text provided for translation.")
			return
		}

		translatesever, err := NewServiceContainer()
		res, err := translatesever.Translator.Translate(text, "en", "zh")
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}
		fmt.Println(res)
	},
}

func init() {
	// 注册子命令到根命令
	rootCmd.AddCommand(entozhCmd)
	// 添加 -m 标志，用于获取翻译的文本
	entozhCmd.Flags().StringP("message", "m", "", "Text to translate")
}
