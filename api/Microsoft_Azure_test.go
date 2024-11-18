package api_test

import (
	"Translate_CLI/api"
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	// 程序输出的结果
	got, err := api.Translate("你好", "zh", "en")
	fmt.Println(got, err)
	// 期望的结果
	want := "Hello"
	if err != nil {
		t.Fatalf("Translate() failed: %v", err)
	}
	// string类型的比较
	if got != want {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestOpenYaml(t *testing.T) {
	got, err := api.OpenYaml()

	// 定义期望的 Config 结构体
	want := api.Config{
		MicrosoftAzure: struct {
			Key      string `yaml:"key"`
			Location string `yaml:"location"`
		}{
			Key:      "12AB34CD56EF789",
			Location: "eastasia",
		},
	}

	// 如果出现错误，测试失败
	if err != nil {
		t.Fatalf("OpenYaml() failed: %v", err)
	}

	// 如果实际结果与期望结果不相等，报告错误
	if got != want {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
