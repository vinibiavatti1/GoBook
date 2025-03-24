// Slice
// Slices are used to store a group of data and have a dynamic size.
// Unlike arrays, slices can be resized dynamically.
// Syntax:
//   []<datatype>{<values>}

package datatypes

import "fmt"

// Declaring Slices
// Demonstrates different ways to declare and initialize slices.
// NOTE: For slices, the Print operation from "fmt" package must be used.
func DeclaringSlices() {

	// Declaring a Slice with Default Values
	slc1 := []int{}
	fmt.Println(slc1) // Output: []

	// Declaring a Slice with Predefined Values
	slc2 := []int{1, 2, 3}
	fmt.Println(slc2) // Output: [1 2 3]
}

// Manipulating Slices
// Demonstrates common operations on slices.
func ManipulatingSlices() {

	// Declaring a Slice
	slc := []string{"A", "B", "C"}

	// Accessing Elements
	el := slc[1]
	fmt.Println(el) // Output: B

	// Getting Length
	ln := len(slc)
	fmt.Println(ln) // Output: 3

	// Adding Elements
	// The built-in "append()" function is used to add elements to a slice dynamically.
	slc = append(slc, "D")
	fmt.Println(slc) // Output: [A Z C D]

	// Adding Multiple Elements
	// We can add more than 1 value using "append()"
	slc = append(slc, "E", "F", "G")
	fmt.Println(slc) // Output: [A Z C D E F G]

	// Resetting Elements
	// We can use the "clear" function to reset all elements of a slice to its default value.
	slc2 := []int{1, 2, 3}
	clear(slc2)
	fmt.Println(slc2) // Output: [0 0 0]

	// Clearing Slice
	// We can clear all elements from a slice by setting the slice using an empty range.
	slc = slc[0:0]
	fmt.Println(slc) // Output: []
}

// Iterating Over Slices
// We can use the for loop to iterate over slices.
func IteratingOverSlices() {

	// Declaring a Slice
	slc := []string{"A", "B", "C"}

	// Iterating over a Slice (Using index-based loop)
	for i := 0; i < len(slc); i++ {
		fmt.Println(i, slc[i]) // Output: 0 A, 1 B, 2 C
	}

	// Iterating over a Slice (Using range)
	for i, v := range slc {
		fmt.Println(i, v) // Output: 0 A, 1 B, 2 C
	}
}

// Unpacking Slices
// We can use the variadic syntax "..." to unpack slices into functions that accept a variable number of arguments.
func UnpackingSlices() {

	// Declaring Slices
	slc1 := []string{"A", "B"}
	slc2 := []string{"C", "D"}

	// Unpacking Slices
	// We can use "..." to unpack slices into variadic functions
	slc3 := append(slc1, slc2...)
	fmt.Println(slc3) // Output: [A B C D]
}
