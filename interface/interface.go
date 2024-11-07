package main

import (
	"errors"
	"fmt"
)

type Printer interface {
	Write()
	// Read()
}

type Moniter struct {
}

func (m Moniter) Write() {
	fmt.Println("Monioter hello")
}

func Print(p Printer) {
	fmt.Println("seocond Monioter hello")

}

type Appliance interface {
	On() error
	Off() error
}

type Laptop struct {
	power bool
}

func (l Laptop) On() error {

	if !l.power {
		return errors.New("need power")
	}
	return nil
}

func (l Laptop) Off() error {
	return nil
}
func main() {

	m := Moniter{}
	m.Write()

	Print(m)
}
