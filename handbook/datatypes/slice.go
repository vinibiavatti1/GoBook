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
	fmt.Println("Default slice:", slc1) // Output: []

	// Declaring a Slice with Predefined Values
	slc2 := []int{1, 2, 3}
	fmt.Println("Predefined slice:", slc2) // Output: [1 2 3]
}

// Manipulating Slices
// Demonstrates common operations on slices.
func ManipulatingSlices() {

	// Declaring a Slice
	slc := []string{"A", "B", "C"}

	// Accessing Elements
	el := slc[1]
	fmt.Println("Accessed element:", el) // Output: B

	// Getting Length
	ln := len(slc)
	fmt.Println("Slice length:", ln) // Output: 3

	// Adding Elements
	// The built-in "append()" function is used to add elements to a slice dynamically.
	slc = append(slc, "D")
	fmt.Println("Slice after append:", slc) // Output: [A Z C D]

	// Adding Multiple Elements
	// We can add more than 1 value using "append()"
	slc = append(slc, "E", "F", "G")
	fmt.Println("Slice after multiple appends:", slc) // Output: [A Z C D E F G]
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
	fmt.Println("Unpacked slice:", slc3) // Output: [A B C D]
}
