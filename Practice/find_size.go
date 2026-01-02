package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var (
		i1 int
		i2 int8
		i3 int16
		i4 int32
		i5 int64
		i6 uint

		f1 float32
		f2 float64

		s1 string

		b1 bool
	)

	// Output in bytes
	fmt.Printf("Type of i1 is %T and its size is %v bytes", i1, unsafe.Sizeof(i1))
	fmt.Printf("\nType of i2 is %T and its size is %v bytes", i2, unsafe.Sizeof(i2))
	fmt.Printf("\nType of i3 is %T and its size is %v bytes", i3, unsafe.Sizeof(i3))
	fmt.Printf("\nType of i4 is %T and its size is %v bytes", i4, unsafe.Sizeof(i4))
	fmt.Printf("\nType of i5 is %T and its size is %v bytes", i5, unsafe.Sizeof(i5))
	fmt.Printf("\nType of i6 is %T and its size is %v bytes", i6, unsafe.Sizeof(i6))

	fmt.Printf("\nType of f1 is %T and its size is %v bytes", f1, unsafe.Sizeof(f1))
	fmt.Printf("\nType of f2 is %T and its size is %v bytes", f2, unsafe.Sizeof(f2))

	fmt.Printf("\nType of s1 is %T and its size is %v bytes", s1, unsafe.Sizeof(s1))

	fmt.Printf("\nType of b1 is %T and its size is %v bytes", b1, unsafe.Sizeof(b1))

	var s string = "vaibhav"
	fmt.Printf("\nThe length of string(number of elements) is %d", len(s))

	fmt.Printf("\nType of s is %v", reflect.TypeOf(s))

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("\nLength - total number of elements of array arr is %v", len(arr))
	fmt.Printf("\nTotal memory size of array arr is %v", unsafe.Sizeof(arr))

}
