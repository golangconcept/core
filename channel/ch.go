package main

import "fmt"

func greet(ch chan string) {
	ch <- "hello this is from channel."
}

func main() {

	ch := make(chan string)

	go greet(ch)

	message := <-ch
	fmt.Println(message)
}
