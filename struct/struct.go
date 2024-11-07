package main

import "fmt"

type laptop struct {
	cpu          string
	ram          int
	storage      int
	manufacturer string
}

func (l laptop) upgradeStorageByValue(size int) {
	l.storage += size
}
func (l *laptop) upgradeStorageByRef(size int) {
	l.storage += size
}
func main() {
	mba := laptop{"M2", 16, 256, "Apple"}

	fmt.Println(mba)

	mba.upgradeStorageByValue(100)
	fmt.Println(mba)
}
