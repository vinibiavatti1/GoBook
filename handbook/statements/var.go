// Variables
// In Go, variables must be declared before they can be used. There are several ways to declare variables:
// 1. Declaring a variable with a specific type
// 2. Declaring a variable using type inference
// 3. Declaring multiple variables in a single line
// NOTE: In Go, an unused variable results in a compilation error

package statements

import "fmt"

// Declaring Variables (Globals)
// These variables are globals and can be accessed by any function.
// inside the package.
var _ int = 1
var _ = 1

// Declaring Multiple Variables (Globals)
// We can use a var block to define multiple variables.
var (
	GlobalVar = 1
	_         = 1 // Unused
)

// Declaring Variables
// There are several ways to declare variables.
func DeclaringVariables() {

	// Declaring Variables
	var x1 int = 10
	fmt.Println(x1) // Output: 10

	// Variables Are Mutable
	x1 = 20
	fmt.Println(x1) // Output: 20

	// Declaring Variables (Type Inference)
	var x2 = 10
	fmt.Println(x2) // Output: 10

	// Declaring Variables (:=)
	x3 := 10
	fmt.Println(x3) // Output: 10

	// Declaring Inline Variables (Inline)
	var a1, b1, c1 int = 1, 2, 3
	fmt.Println(a1, b1, c1) // Output: 1 2 3

	// Declaring Inline Variables (Type Inference)
	var a2, b2, c2 = 1, true, 3.5
	fmt.Println(a2, b2, c2) // Output: 1 true 3.5

	// Declaring Inline Variables (:=)
	a3, b3, c3 := 1, true, 3.5
	fmt.Println(a3, b3, c3) // Output: 1 true 3.5

	// Declaring Multiple Variables
	var (
		a4 = 1
		b4 = true
		c4 = "Hello"
	)
	fmt.Println(a4, b4, c4) // Output: 1 true Hello World!

	// Declaring Unused Variables
	var _ int = 1

	// Declaring Unused Variables (Type Inference)
	var _ = 1
}

// Setting Variables
// Variables are mutable.
func SettingVariables() {

	// Declaring Variables
	var x = 1

	// Changing Variable Values
	x = 2
	fmt.Println(x) // Output: 2
}
