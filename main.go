package main

import (
	"fmt"
	"slices"
	"unsafe"
)

func main() {
	fmt.Println("This is a bool 2 int benchmark thingy")
}

func bool2intSimplest(b bool) int {
	if b {
		return 1
	}
	return 0
}

// @see https://0x0f.me/blog/golang-compiler-optimization/
func bool2intFastest(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func bool2intSwitch(b bool) int {
	switch b {
	case true:
		return 1
	case false:
		return 0
	default:
		panic("this doesn't really need to exist")
	}
}

func bool2intInternalMap(b bool) int {
	table := map[bool]int{
		true:  1,
		false: 0,
	}
	return table[b]
}

var table = map[bool]int{
	true:  1,
	false: 0,
}

func bool2intExternalMap(b bool) int {
	return table[b]
}

var list = []bool{
	false,
	true,
}

func bool2intSlice(b bool) int {
	return slices.Index(list, b)
}

func bool2intUnsafePointer(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}
