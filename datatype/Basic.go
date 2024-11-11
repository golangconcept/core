package main

import (
	"fmt"
	"unsafe"
)

func main() {
	j := 3.4
	k := j
	l := int32(4)
	fmt.Println(j, k, l, unsafe.Sizeof(l), l.(type))
}
