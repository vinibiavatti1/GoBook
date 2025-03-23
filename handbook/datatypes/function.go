// Function
// Functions in Go are defined using the "func" keyword.
// They can have parameters, return values, and support variadic arguments.
// Syntax:
//   func <name>(<parameters> <type>...) <return> {...}

package datatypes

import "fmt"

// Void Function
// A function that does not return values.
func Func() {
}

// Function With Parameters
// A function that takes two integer parameters and prints their sum.
func FuncWithParams(x int, y int) {
	fmt.Println(x + y)
}

// Function With Grouped Parameters
// We can omit repeated parameter types when the types are the same.
func FuncWithGroupedParams(x, y int) {
	fmt.Println(x + y)
}

// Func With Return
// A function that takes two integers and returns their sum.
func FuncWithReturn(x, y int) int {
	return x + y
}

// Func With Multiple Return Values
// A function that returns both the sum and product of two integers.
func FuncWithMultipleReturn(x, y int) (int, int) {
	return x + y, x * y
}

// Named Return Values
// We can define named return values directly in the function signature.
// This allows the function to return values implicitly without explicitly specifying them in the return statement.
func NamedReturnValues() (x, y int) {
	x, y = 1, 1 // Assigning to the named return variable
	return      // Implicitly returns "x" and "y"
}

// Variadic Function
// A function that accepts a variable number of arguments.
// Usage examples: Fun4(1, 2) Fun4(1, 2, 3, 4)
func VariadicFunc(nums ...int) {
	fmt.Println(nums) // Prints the slice containing all arguments
}

// Function As Parameter
// Functions can be passed as arguments to other functions.
// This allows for higher-order functions and more flexible behavior.
func FuncAsParameter(fn func()) {
	fn()
}

// Function As Return Type
// Functions in Go can also return other functions, enabling closures and deferred execution.
func FuncAsReturnType() func() {
	return func() {}
}

// Unused Func
// A blank identifier function.
func _() {
}

// Assigning Functions
// We can assign a lambda function to a variable and use that variable
// as a function reference.
func LambdaFunc() {

	// A Simple Lambda Function
	fn1 := func() {
		fmt.Println("Hello, World!")
	}
	fn1() // Output: Hello, World!

	// A Lambda Function With Parameters and Return Value
	fn2 := func(y int) int {
		return y * 2
	}
	fmt.Println(fn2(5)) // Output: 10

	// Immediately Invoked Function Expression (IIFE)
	// We define and call the function at the same time.
	x := func(y int) int {
		return y * 2
	}(5)
	fmt.Println(x) // Output: 10
}

// Closure Function
// A closure is a lambda function returned by another function.
// It can maintain the state of variables across multiple calls.
func Closure() func() int {
	x := -1
	return func() int {
		x++ // Increment x and return the updated value
		return x
	}
}

// Closure Usage
// Demonstrating how closures retain state across multiple invocations.
func ClosureUsage() {

	// Creating a new closure instance
	counter := Closure()

	// Calling the closure multiple times
	fmt.Println(counter()) // Output: 0
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
}
