package main

import "fmt"

func main() {

	// ===================================================
	fmt.Println("=========================================")
	fmt.Println("======== Array with fixed length ========")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with fixed length:", arr)

	// arr = append(arr, 6) // This will throw an error because arrays are fixed length

	fmt.Println("=========================================")
	fmt.Println("========= Slice with zero length ========")

	slice := make([]int, 0, 10) // len(0) cap(10)

	fmt.Println("Slice with zero length and ten capacity:", slice)
	fmt.Println("Length of the Slice:", len(slice))
	fmt.Println("Capacity of the Slice:", cap(slice))

	slice = append(slice, 99)

	fmt.Println("Slice after appending 99:", slice)
}
