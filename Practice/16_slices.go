package main

import "fmt"

func main() {
	grades := []int{1, 2, 3}
	fmt.Printf("Grades: %d\n", grades)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("My array: %d\n", arr)
	fmt.Printf("Length of array: %d\n", len(arr))
	fmt.Printf("Capacity of array: %d\n", cap(arr))

	slice_1 := arr[1:4]
	fmt.Printf("Slice 1: %d\n", slice_1)
	fmt.Printf("Length of slice_1: %d\n", len(slice_1))
	fmt.Printf("Capacity of slice_1: %d\n", cap(slice_1))

	sub_slice := slice_1[1:3]
	fmt.Printf("Sub slice: %d\n", sub_slice)

	// using make function

	slice_2 := make([]int, 5, 8)
	fmt.Printf("slice_2: %d\n", slice_2)
	fmt.Printf("length of slice_2: %d\n", len(slice_2))
	fmt.Printf("capacity of slice_2: %d\n", cap(slice_2))

	arr_3 := [5]int{10, 20, 30, 40, 50}
	slice_3 := arr_3[:3]

	fmt.Printf("Array_3: %d\n", arr_3)
	fmt.Printf("Slice_3: %d\n", slice_3)

	slice_3[1] = 200
	fmt.Printf("Array_3: %d\n", arr_3)
	fmt.Printf("Slice_3: %d\n", slice_3)

	myarr := [4]int{10, 20, 30, 40}
	fmt.Printf("My array: %d, legth: %d, capacity: %d\n", myarr, len(myarr), cap(myarr))

	myslice := myarr[1:3]
	myslice_1 := append(myslice, 900)
	myslice_2 := append(myslice, 90, 100)

	fmt.Printf("My array: %d, legth: %d, capacity: %d\n", myarr, len(myarr), cap(myarr))
	fmt.Printf("My slice: %d, legth: %d, capacity: %d\n", myslice, len(myslice), cap(myslice))
	fmt.Printf("My slice_1: %d, legth: %d, capacity: %d\n", myslice_1, len(myslice_1), cap(myslice_1))
	fmt.Printf("My slice_2: %d, legth: %d, capacity: %d\n", myslice_2, len(myslice_2), cap(myslice_2))

	mynew_slice := append(myslice_1, myslice_2...)
	fmt.Printf("My new slice: %d, legth: %d, capacity: %d\n", mynew_slice, len(mynew_slice), cap(mynew_slice))

	// deleting from slice
	delarr := [5]int{10, 20, 30, 40, 50}
	i := 2 // index to delete
	fmt.Printf("Array: %d\n", delarr)
	delslice_1 := delarr[:i]
	delslice_2 := delarr[i+1:]
	delslice := append(delslice_1, delslice_2...)
	fmt.Printf("Slice: %d\n", delslice)

	// copy from slice
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slic := make([]int, 3)
	num := copy(dest_slic, src_slice) // returns number of elements copied
	fmt.Printf("Number of elements copied: %d\n", num)
	fmt.Printf("Destination slice: %d\n", dest_slic)
}
