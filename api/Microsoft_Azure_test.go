package api_test

import (
	"Translate_CLI/cmd"
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	// 程序输出的结果
	translatesever, err := cmd.NewServiceContainer()
	got, err := translatesever.Translator.Translate("你好", "zh", "en")
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

func TestTranslate2(t *testing.T) {
	// 程序输出的结果
	translatesever, err := cmd.NewServiceContainer()
	got, err := translatesever.Translator.Translate("hello hello", "en", "zh")
	fmt.Println(got, err)
	// 期望的结果
	want := "你好 好的"
	if err != nil {
		t.Fatalf("Translate() failed: %v", err)
	}
	// string类型的比较
	if got != want {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
