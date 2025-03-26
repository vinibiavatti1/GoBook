// Slices
// Slices are used to store a sequence of elements of the same type.
// Unlike arrays, slices can be resized dynamically.
// Syntax:
//   []<datatype>{<values>}

package composite

import "fmt"

// Declaring Slices
// Slices are declared without specifying a length.
// If a length is specified, an array will be created instead.
func DeclaringSlices() {

	// Declaring a Slice
	// When no elements are specified, the slice will be empty.
	x := []int{}
	fmt.Println("x:", x) // Output: x: []

	// Declaring a Slice with Predefined Values
	// We can initialize the slice with predefined values.
	x = []int{1, 2, 3}
	fmt.Println("x:", x) // Output: [1 2 3]
}

// Manipulating Slices
// We can add, access, reset and clear elements in a slice.
func ManipulatingSlices() {

	// Declaring a Slice
	// We will declare a slice with 3 elements to be used in the examples.
	x := []string{"A", "B", "C"}

	// Accessing Elements
	// We can access elements using indexes [i].
	y := x[1]
	fmt.Println("y:", y) // Output: y: B

	// Slicing Slices
	// We can slice slices using the [start:end] syntax.
	// Note: The end index is exclusive.
	// Syntax: x[<start>:<end>]
	fmt.Println("x[ :5]", x[:2])  // Output: [A B]
	fmt.Println("x[7: ]", x[1:])  // Output: [B C]
	fmt.Println("x[2:4]", x[1:2]) // Output: [B]
	fmt.Println("x[ : ]", x[:])   // Output: [A B C] (Same reference)
	fmt.Println("x[0:0]", x[0:0]) // Output: [] (empty array)

	// Retrieving Length
	// We can get the length of a slice using the "len()" builtin function.
	l := len(x)
	fmt.Println("len:", l) // Output: len: 3

	// Mutating Elements
	// We can mutate elements using indexes [i].
	x[1] = "Z"
	fmt.Println("x:", x) // Output: x: [A Z C]

	// Appending Elements
	// We can use the "append()" builtin function to add elements to a slice.
	x = append(x, "D")
	fmt.Println("x:", x) // Output: x: [A Z C D]

	// Appending Multiple Elements
	// The "append()" function accepts a variadic number of elements.
	x = append(x, "E", "F", "G")
	fmt.Println("x:", x) // Output: [A Z C D E F G]

	// Resetting Elements
	// We can use the "clear" function to reset all elements of a slice to its default value.
	a := []int{1, 2, 3}
	clear(a)
	fmt.Println("a:", a) // Output: a: [0 0 0]

	// Clearing Slice
	// We can clear all elements from a slice by setting the slice using an empty range.
	a = a[0:0]
	fmt.Println("a:", a) // Output: x: []

	// Unpacking Slices
	// We can use the spread operator "..." to unpack slices into functions that accept a variable number of arguments.
	a = append(a, []int{1, 2, 3}...)
	fmt.Println("a:", a) // Output: a: [1 2 3]

	// Iterating Slices (Using standard for)
	// We can use the standard for loop to iterate over a slice
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i]) // Output: 0 A, 1 Z, 2 C, 3 D, 4 E, 5 F, 6 G
	}

	// Iterating Slices (Using for-range)
	// We can use the for-range loop to iterate over a slice
	for i, v := range x {
		fmt.Println(i, v) // Output: 0 A, 1 Z, 2 C, 3 D, 4 E, 5 F, 6 G
	}
}
