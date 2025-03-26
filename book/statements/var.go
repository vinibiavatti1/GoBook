// Variables
// In Go, variables must be declared before they can be used. There are several ways to declare variables:
// - Declaring a variable with a specific type
// - Declaring a variable using type inference
// - Declaring multiple variables in a single line
// Variables can be global or local.
// Global variables can be accessed by any function inside the package.
// Note: In Go, an unused variable results in a compilation error

package statements

import "fmt"

// Declaring Global Variables
// We can declare global variables outside of functions.
// Global variables can be accessed by any function inside the package.
// The ":=" operator cannot be used to declare global variables.
var Var1 int = 1      // Type explicitly
var Var2 = 2          // Type Inference
var Var3, Var4 = 3, 4 // Inline
var _ int = 1         // Unused

// Declaring Multiple Global Variables
// We can declare multiple global variables in multiple lines using the var keyword.
var (
	Var5       int = 1    // Type explicitly
	Var6           = 2    // Type Inference
	Var7, Var8     = 3, 4 // Inline
	_              = 1    // Unused
)

// Declaring Variables
// We can declare local variables inside functions.
// Variables must be declared before they can be used.
func DeclaringVariables() {

	// Declaring a Variable
	// In this example, we declare a variable with the type int explicitly.
	var x1 int = 10
	fmt.Println("x1:", x1) // Output: x1: 10

	// Declaring a Variable with Type Inference
	// The compiler can infer the type of a variable based on the value assigned to it.
	var x2 = 10
	fmt.Println("x2:", x2) // Output: x2: 10

	// Declaring a Variable with Short Declaration Operator (:=)
	// The short declaration operator (:=) is used to declare and initialize a variable.
	x3 := 10
	fmt.Println("x3:", x3) // Output: x3: 10

	// Declaring Unused Variables
	// We can use the blank identifier (_) to declare unused variables.
	// The short declaration operator (:=) cannot be used to declare unused variables.
	var _ int = 1
	var _ = 1

	// Variables are Mutable
	// Different from constants, variables are mutable.
	// We can change the value of a variable after it has been declared.
	x1 = 20
	fmt.Println("x1:", x1) // Output: x1: 20 (old value: 10)

	// Declaring Inline Variables
	// We can declare multiple variables in a single line.
	// Note that the last type declared is used for all variables.
	var a1, b1, c1 int = 1, 2, 3
	fmt.Println("Inline:", a1, b1, c1) // Output: Inline: 1 2 3

	// Declaring Inline Variables With Type Inference
	// We can declare multiple variables in a single line using type inference.
	var a2, b2, c2 = 1, true, 3.5
	fmt.Println("Inline:", a2, b2, c2) // Output: Inline: 1 true 3.5

	// Declaring Inline Variables With Short Declaration Operator (:=)
	// We can declare and initialize multiple variables in a single line using the short declaration operator (:=).
	a3, b3, c3 := 1, true, 3.5
	fmt.Println("Inline:", a3, b3, c3) // Output: Inline: 1 true 3.5

	// Declaring Multiple Variables
	// We can declare multiple variables in multiple lines using the var keyword.
	var (
		var1       int = 1    // Type explicitly
		var2           = 2    // Type Inference
		var3, var4     = 3, 4 // Inline
		_              = 5    // Unused
	)
	fmt.Println("Vars:", var1, var2, var3, var4) // Output: Vars: 1 2 3 4
}
