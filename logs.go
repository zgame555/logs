package logs

import "fmt"

func Info(param ...any) {
	param = append([]any{"Info"}, param...)
	fmt.Println(param)
}

func Error(param ...any) {
	param = append([]any{"Error"}, param...)
	fmt.Println(param)
}

func Warn(param ...any) {
	param = append([]any{"Warn"}, param...)
	fmt.Println(param)
}
