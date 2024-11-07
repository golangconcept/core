package main

import "fmt"

type Printer interface {
	Write()
	Read()
}

type Moniter struct {
}

func (m Moniter) Write() {
	fmt.Println("Monioter hello")
}

func main() {

	m := Moniter{}
	m.Write()
}
