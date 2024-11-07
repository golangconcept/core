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
func (l *laptop) getCpu() string {
	return l.cpu
}
func (l *laptop) setCpu(newCpu string) {
	l.cpu = newCpu
}
func main() {
	mba := laptop{"M2", 16, 256, "Apple"}

	fmt.Println(mba)

	mba.upgradeStorageByValue(100)
	mba.setCpu("M3")
	fmt.Println(mba.getCpu())

	mba.upgradeStorageByRef(200)
	fmt.Println(mba)
}
