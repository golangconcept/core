package main

import (
	"fmt"
	"runtime"
)

func rec() {
	if h := recover(); h != nil {

		_, file, line, _ := runtime.Caller(2)
		fmt.Printf("Panic occurred at %s : %d\n", file, line)
		fmt.Println(h)
	}
}
func p1() {
	panic("this is panic 2")
}
func main() {
	defer rec()
	panic("this is panic")
	// p1()

}
