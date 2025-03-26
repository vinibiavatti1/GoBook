// If
// The "if" keyword is used to execute code based on conditions.
// It allows you to control the flow of execution by evaluating conditions.
// Syntax:
//   if <condition> {...}
//   if <condition> {...} else {...}
//   if <condition> {...} else if {...}

package statements

import "fmt"

// If with Comparison Operators
// You can use comparison operators (e.g., ==, !=, >, <) in if conditions.
func IfWithComparisonOperators() {

	// Declaring Variables
	x, y := 1, 1

	// Equals Condition
	if x == y {
		fmt.Println("x == y")
	}

	// Not Equals Condition
	if x != y {
		fmt.Println("x != y")
	}

	// Greater Than Condition
	if x > y {
		fmt.Println("x > y")
	}

	// Greater Than Or Equal To Condition
	if x >= y {
		fmt.Println("x >= y")
	}

	// Less Than Condition
	if x < y {
		fmt.Println("x < y")
	}

	// Less Than Or Equal To Condition
	if x <= y {
		fmt.Println("x <= y")
	}
}

// If With Logical Operators (&&, || e !)
// Logical operators can be used in if conditions to combine multiple conditions.
func IfWithLogicalOperators() {

	// Declaring Variables
	x, y, z := 1, 1, 1

	// AND (&&)
	if x == y && x == z {
		fmt.Println("x == y && x == z")
	}

	// OR (||)
	if x == y || x == z {
		fmt.Println("x == y || x == z")
	}

	// NOT (!)
	if !(x == z) {
		fmt.Println("x != z")
	}
}

// If Else
// The else block is performed when the condition is false
// We can create chained conditions with Else-If
func IfElse() {

	// Declaring Variables
	x, y, z := 1, 1, 1

	// If-Else
	if x == y {
		fmt.Println("x == y") // Output: x == y
	} else {
		fmt.Println("x != y")
	}

	// Else-If
	if x == y {
		fmt.Println("x == y") // Output: x == y
	} else if x == z {
		fmt.Println("x == z")
	}

	// Else-If (With Else)
	if x == y {
		fmt.Println("x == y") // Output: x == y
	} else if x == z {
		fmt.Println("x == z")
	} else {
		fmt.Println("x != y && x != z")
	}

	// Else-If (With Switch)
	// A switch statement without an expression have the same behavior of a Else-If
	switch {
	case x == y:
		fmt.Println("x == y") // Output: x == y
	case x == z:
		fmt.Println("x == z")
	default:
		fmt.Println("x != y && x != z")
	}
}

// If With Declaration
// We can declare variables inside the if condition.
func IfWIthDeclaration() {

	// If with Declaration
	// We can declare a variable inside the if condition.
	if x := true; x {
		fmt.Println("x is true")
	}

	// If with multiple declarations
	// We can declare multiple variables inside the if condition.
	if x, y := 1, 2; x < y {
		fmt.Println("x is less than y")
	}
}
