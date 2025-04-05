package main

import (
	"fmt"
)

// func rec() {
// 	if h := recover(); h != nil {

//			_, file, line, _ := runtime.Caller(2)
//			fmt.Printf("Panic occurred at %s : %d\n", file, line)
//			fmt.Println(h)
//		}
//	}
// package main

func causePanic1() {
	panic("First panic!")
}

func causePanic2() string {
	panic("Second panic!")
}

func handlePanics() (result string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
			result = "Recovered from panic situaltions."
		}
	}()

	causePanic1() // This will panic
	// causePanic2() // This will also panic, but we can recover from it
	result = causePanic2()
	return result
}

func main() {
	handlePanics()
	fmt.Println("Program continues without crashing")
}
