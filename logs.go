package logs

import "fmt"

func Info(param ...any) {
	fmt.Printf("Info : %v\n", param)
}

func Error(param ...any) {
	fmt.Printf("Error : %v\n", param)
}

func Warn(param ...any) {
	fmt.Printf("Warn : %v\n", param)
}
