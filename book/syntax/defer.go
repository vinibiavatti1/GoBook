// Defer
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed
// until the surrounding function returns.
// The execution order of deferred functions is LIFO (Last In First Out).
// Deferred functions avoid needing to write closing syntax, such as closing a file or
// unlocking a mutex for each code case.

package syntax

import "fmt"

// Single Defer
// A defer statement defers the execution of a function until the surrounding function returns.
// This exmaple shows a single defer statement.
// The End message is printed before the Deferred message.
func singleDefer() {
	defer fmt.Println("Defer")
	fmt.Println("End")
}
func PerformDefer() {
	singleDefer()
	// Output: End
	// Output: Defer
}

// Multiple Defers
// The deferred functions are executed in reverse order.
// The execution order of deferred functions is LIFO (Last In First Out).
func multipleDefers() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println("End")
}
func PerformMultipleDefers() {
	multipleDefers()
	// Output: End
	// Output: Defer 2
	// Output: Defer 1
}

// Defer Arguments
// The deferred call's arguments are evaluated immediately, but the function call is not executed
// until the surrounding function returns.
// Notice that the deferred function prints "Deferred" instead of "End".
func deferArguments() {
	x := "Deferred"
	defer fmt.Println(x)
	x = "End"
}
func PerformDeferArguments() {
	deferArguments()
	// Output: Deferred
}

// Defer Loop
// The deferred function is executed after the loop ends.
// As we can see, the values of i are printed in reverse order.
func deferLoop() {
	for i := range 3 {
		defer fmt.Println(i)
	}
}
func PerformDeferLoop() {
	deferLoop()
	// Output: 2
	// Output: 1
	// Output: 0
}

// Defer with Function
// We can defer a function execution.
// The deferred function is executed after the surrounding function returns.
// Notice that the function must be Immediately Invoked.
func deferFunction() {
	defer func() { fmt.Println("Function") }()
	fmt.Println("End")
}
func PerformDeferFunction() {
	deferFunction()
	// Output: End
	// Output: Function
}

// Defer Named Return
// A deferred function can modify the return value of a named return.
// The deferred function modifies the return value of x.
func deferNamedReturn() (x int) {
	x = 5
	defer func() { x += 10 }()
	return
}
func PerformDeferNamedReturn() {
	fmt.Println(deferNamedReturn())
	// Output: 15
}
