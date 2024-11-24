package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

var defaultCmd = &cobra.Command{
	Use:   "d", // 子命令名称
	Short: "Default translation options",
	Long: `If you enter English, it will be translated into Chinese,
and if you enter Chinese, it will be translated into English,
other translation options are not supported at this time`,
	Args: cobra.ExactArgs(1), // 确保必须传入一个参数
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0] // 从命令行参数获取输入

		targetLang := "zh"
		sourceLang := "en"
		if containsChinese(text) {
			targetLang = "en" // 如果包含中文，翻译目标语言为英文
			sourceLang = "zh"
		}

		translatesever, err := NewServiceContainer()

		res, err := translatesever.Translator.Translate(text, sourceLang, targetLang)
		if err != nil {
			fmt.Printf("error:%v", err)
			return
		}
		fmt.Println(res)
	},
}

func containsChinese(s string) bool {
	var hzRegexp, _ = regexp.Compile("([\u4e00-\u9fa5]+)")
	return hzRegexp.MatchString(s)
}

func init() {
	// 注册子命令到根命令
	rootCmd.AddCommand(defaultCmd)
}
