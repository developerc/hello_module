package hellomodule

import "fmt"

// Pretty печатает переданные аргументы
func Pretty(v ...interface{}) {
	fmt.Println(v...)
}
