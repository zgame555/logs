package logs

import "fmt"

func Info(param ...any) {
	param = append([]any{"✅"}, param...)
	fmt.Println(param...)
}

func Error(param ...any) {
	param = append([]any{"❌"}, param...)
	fmt.Println(param...)
}

func Warn(param ...any) {
	param = append([]any{"⚠️"}, param...)
	fmt.Println(param...)
}
