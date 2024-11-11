package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	j := 3.4
	k := j
	l := int32(4)
	fmt.Println(j, k, l, unsafe.Sizeof(l), reflect.TypeOf(l), reflect.TypeOf(l).Kind())
}
