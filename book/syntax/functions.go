// Functions
// Functions in Go are defined using the "func" keyword.
// They can take zero or more parameters and return zero or more values.
// Syntax:
//   func <name>(<parameters> <type>...) <return> {...}

package syntax

import "fmt"

// Private Functions
// Functions with the name starting with a lowercase letter are private functions.
// They can only be called inside the package where they are defined.
func privateFunction() {
}

// Public Functions
// Functions with the name starting with an uppercase letter are public functions.
// They can be called from other packages.
func PublicFunction() {
	privateFunction() // Call the private function example
}

// Parameters
// Functions can take zero or more parameters.
// Parameters are defined with a name and a type.
func FunctionParameters(x int, y string) {
}

// Grouped Parameters
// Multiple parameters of the same type can be grouped together.
func GroupedParameters(x, y, z int) {
}

// Return Value
// Functions can return zero or more values.
// Return values are defined after the parameter list.
func FunctionReturn() int {
	return 1
}

// Multiple Return Values
// Functions can return multiple values.
// The return values are separated by commas and enclosed in parentheses.
func FunctionMultipleReturn() (int, int) {
	return 1, 2
}

// Named Return Values
// Functions can return named values.
// Named return values are initialized with the zero value of their type.
// An empty return statement returns the named values implicitly.
func NamedReturnValues() (x, y int) {
	x, y = 1, 1 // Assigning to the named return variable
	return      // Implicitly returns "x" and "y"
}

// Variadic Parameters
// Functions can take a variable number of arguments using variadic parameters.
// The variadic parameter is a slice that contains all arguments passed to the function.
// To define a variadic parameter, use the "..." operator followed by the type.
func VariadicParameters(nums ...int) {
	fmt.Println(nums) // Output: [1 2 3]
}

// Function As Parameter
// Functions can be passed as parameters to other functions.
// This is useful for callbacks and deferred execution.
func FunctionAsParameter(fn func()) {
	fn()
}

// Function As Return Type
// Functions can be returned as values from other functions.
// This is useful for creating closures and higher-order functions.
func FunctionAsReturnType() func() {
	return func() {}
}

// Unused Function
// Unused functions can be declared using the blank identifier "_".
func _() {
}

// Lambda Functions
// Lambda functions are anonymous functions that can be assigned to variables.
// They are useful for creating closures and deferred execution.
func LambdaFunctions() {

	// Simple Lambda Function
	// An example of a lambda function without parameters and return values.
	fn1 := func() {
		fmt.Println("Hello, World!")
	}
	fn1() // Output: Hello, World!

	// Lambda Function with Parameters and Return Value
	// An example of a lambda function with parameters and a return values.
	fn2 := func(x int) int {
		return x * 2
	}
	x := fn2(2)
	fmt.Println("x:", x) // Output: x: 4

	// Immediately Invoked Function Expression (IIFE)
	// An example of an immediately invoked lambda function.
	// This is useful for deferred execution.
	x = func(y int) int {
		return y * 2
	}(3)
	fmt.Println("x:", x) // Output: 6
}

// Closure Functions
// Closures are functions that capture variables from their surrounding context.
// They can keep state across multiple invocations.
func ReturnClosure() func() int {
	x := -1 // Variable captured by the closure
	return func() int {
		x++ // Increment x and return the updated value
		return x
	}
}

// Using Closures
// In this example, we create a closure that increments a counter.
// Note that the closure keeps the state of the counter across multiple invocations.
func UsingClosures() {

	// Call the Function
	// The function returns a closure that increments a counter.
	next := ReturnClosure()

	// Calling the Closure
	// Note that the closure keeps the state of the counter across multiple invocations.
	fmt.Println("Counter:", next()) // Output: Counter: 0
	fmt.Println("Counter:", next()) // Output: Counter: 1
	fmt.Println("Counter:", next()) // Output: Counter: 2
}
