// Arrays
// Arrays are a collection of elements of the same type.
// They are fixed in size and have a zero-based index.
// Syntax:
//   [<length>]<datatype>{<values>}

package composite

import "fmt"

// Declaring Arrays
// Arrays must be declared with a fixed length and type.
// If the length is not specified, a slice will be created instead.
func DeclaringArrays() {

	// Declaring an Array
	// All elements are initialized with the zero value of the type.
	x := [3]int{}
	fmt.Println("x:", x) // Output: x: [0 0 0]

	// Declaring an Array with Predefined Values
	// We can initialize the array with predefined values.
	x = [3]int{1, 2, 3}
	fmt.Println("x:", x) // Output: x: [1 2 3]

	// Declaring an Array with Inferred Length
	// We can use the ellipsis operator "..." to infer the length of the array based on the
	// number of elements.
	x = [...]int{1, 2, 3}
	fmt.Println("x:", x, "len:", len(x)) // Output: x: [1 2 3] len: 3
}

// Manipulating Arrays
// We can access, mutate, and get the length of an array.
func ManipulatingArrays() {

	// Declaring an Array
	// We will declare an array with 3 elements to be used in the examples.
	x := [3]string{"A", "B", "C"}

	// Acessing Elements
	// We can access elements using indexes [i].
	y := x[1]
	fmt.Println("Element:", y) // Output: Element: B

	// Slicing Arrays
	// We can slice arrays using the [start:end] syntax.
	// Note: The end index is exclusive.
	// Syntax: x[<start>:<end>]
	fmt.Println("x[ :5]", x[:2])  // Output: [A B]
	fmt.Println("x[7: ]", x[1:])  // Output: [B C]
	fmt.Println("x[2:4]", x[1:2]) // Output: [B]
	fmt.Println("x[ : ]", x[:])   // Output: [A B C] (Same reference)
	fmt.Println("x[0:0]", x[0:0]) // Output: [] (empty array)

	// Retrieving Length
	// We can get the length of an array using the "len()" builtin function.
	l := len(x)
	fmt.Println("len:", l) // Output: 3

	// Mutating Elements
	// We can mutate elements using indexes [i].
	// Note: Arrays are fixed in size, so we cannot append or remove elements.
	x[1] = "Z"
	fmt.Println("x:", x) // Output: x: [A Z C]

	// Iterating Arrays (Using standard for)
	// We can use the standard for loop to iterate over an array.
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i]) // Output: 0 A, 1 Z, 2 C
	}

	// Iterating Arrays (Using for-range)
	// We can use the for-range loop to iterate over an array.
	for i, v := range x {
		fmt.Println(i, v) // Output: 0 A, 1 Z, 2 C
	}
}
