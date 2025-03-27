// ============= ####### =============
// ============== main ==============
// ============= ####### =============

package main

import "fmt"

// ============= ####### =============
// ============== main ==============
// ============= ####### =============

func main() {
	// ============= ####### =============
	// ======== fixed length array ======
	// ============= ####### =============

	// Create an array with fixed length of 5
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with fixed length:", arr)

	// Note: Arrays are fixed length
	// The following would cause an error:
	// arr = append(arr, 6)

	// ============= ####### =============
	// ============= slices =============
	// ============= ####### =============

	// Create a slice with length 0 and capacity 10
	slice := make([]int, 0, 10) // len(0) cap(10)

	// Demonstrate slice properties
	fmt.Println("Slice with zero length and ten capacity:", slice)
	fmt.Println("Length of the Slice:", len(slice))
	fmt.Println("Capacity of the Slice:", cap(slice))

	// Demonstrate slice growth
	slice = append(slice, 99)
	fmt.Println("Slice after appending 99:", slice)
}
