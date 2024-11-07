package main

import "fmt"

type laptop struct {
	cpu          string
	ram          int
	storage      int
	manufacturer string
}

func main() {
	mba := laptop{"M2", 16, 256, "Apple"}

	fmt.Println(mba)
}
