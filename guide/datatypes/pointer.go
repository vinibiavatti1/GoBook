// Pointer
// Pointers are a way to reference a variable's memory address in Go.
// They allow you to modify the value of a variable without passing it by value.
// Pointers are declared using the `*` operator, and you can obtain a pointer to a variable using the `&` operator.
// Syntax:
//   *x - dereference operator, used to access the value at the address stored in x
//   &x - address operator, used to get the memory address of x

package datatypes

import "fmt"

// Using Pointers
// The function below demonstrates how to use pointers in Go.
func UsingPointers() {

	// Declaring a Variable
	// We will declare a variable to use as example.
	x := 5
	fmt.Println("x:", x) // Output: x: 5

	// Declaring a Pointer
	// We will declare a pointer to hold the address of x.
	// We can use the "&" operator to get the address of x.
	p := &x
	fmt.Println("p:", p) // Output: p: 0xc00000a0b8

	// Dereferencing a Pointer
	// We can use the "*" operator to dereference the pointer and get the value at the address.
	y := *p
	fmt.Println("y:", y) // Output: y: 5

	// Changing the Value of a Pointer
	// Functions can receive pointers instead of values, to manipulate the original variable.
	changeValue(p)
	fmt.Println("x:", *p) // Output: x: 10

	// Pointers Can Have Nil Values
	// A pointer can be nil, meaning it doesn't point to any address.
	p = nil
	fmt.Println("p:", p) // Output: p: <nil>
}

// Change Value
// Since the x parameter is a pointer, the value of the argument will not be copied.
// Instead, the reference will be passed to the function, allowing us to modify the original variable.
func changeValue(x *int) {
	*x = 10
}
