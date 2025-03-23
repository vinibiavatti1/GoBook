// Arrays
// Arrays are used to store an group of data, and have a fixed size.
// Unlike slices, arrays cannot be resized dynamically.
// Syntax
//   [<length>]<datatype>{<values>}

package datatypes

import "fmt"

// Declaring Arrays
// Demonstrates different ways to declare and initialize arrays.
func DeclaringArrays() {

	// Declaring an Array with Default Values
	arr1 := [3]int{}
	fmt.Println(arr1) // Output: [0 0 0]

	// Declaring an Array with Predefined Values
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2) // Output: [1 2 3]

	// Declaring an Array with Inferred Length
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3) // Output: [1 2 3]
}

// Manipulating Arrays
// Demonstrates common operations on arrays.
func ManipulatingArrays() {

	// Declaring an Array
	arr := [3]string{"A", "B", "C"}

	// Acessing Elements
	el := arr[1]
	fmt.Println(el) // Output: B

	// Getting Length
	ln := len(arr)
	fmt.Println(ln) // Output: 3

	// Mutating Elements
	arr[1] = "Z"
	fmt.Println(arr) // Output: [A Z C]
}

// Iterating Over Arrays
// We can use the for operation to iterating over arrays
func IteratingOverArrays() {

	// Declaring an Array
	arr := [3]string{"A", "B", "C"}

	// Iterating over an Array (Using for index)
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i]) // Output: 0 A, 1 B, 2 C
	}

	// Iterating over an Array (Using range)
	for i, v := range arr {
		fmt.Println(i, v) // Output: 0 A, 1 B, 2 C
	}
}
