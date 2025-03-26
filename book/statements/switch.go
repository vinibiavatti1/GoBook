// Switch
// The "switch" statement is used to execute different blocks of code based on the value of an expression.
// It provides a cleaner and more readable alternative to multiple "if-else" conditions.
// Syntax
//   switch <var> {
//   case <expr>, <expr>:
//       ...
//       fallthrough
//   default:
//       ...
//   }

package statements

import "fmt"

func Switches() {

	// Declaring a Variable
	x := 2

	// Basic Switch
	// A standard switch statement evaluates "x" and executes the matching case.
	// Unlike other languages, Go does not require a "break" statement to prevent fallthrough.
	switch x {
	case 1:
		fmt.Println("x == 1")
	case 2:
		fmt.Println("x == 2") // (Matched) Output: x == 2
	}

	// Switch with Default Case
	// The "default" case executes when no other case matches.
	switch x {
	case 1:
		fmt.Println("x == 1")
	default:
		fmt.Println("?") // (Matched) Output: ?
	}

	// Case with Multiple Values
	// A single case can match multiple values, acting as an OR condition.
	switch x {
	case 1, 2:
		fmt.Println("x == 1 || x == 2") // (Matched) Output: x == 1 || x == 2
	default:
		fmt.Println("?")
	}

	// Switch with Expressions
	// Cases can evaluate expressions dynamically, similar to "if" conditions.
	y := 1
	switch x {
	case y - 1:
		fmt.Println("x == y - 1")
	case y + 1:
		fmt.Println("x == y + 1") // (Matched) Output: x == y + 1
	}

	// Type Switch
	// A type switch is used to determine the type of an interface{} (or "any" in Go 1.18+).
	var z any = true
	switch z.(type) {
	case int:
		fmt.Println("z is an int")
	case bool:
		fmt.Println("z is a bool") // (Matched) Output: z is a bool
	}

	// Fallthrough
	// The "fallthrough" keyword forces execution of the next case, even if it does not match.
	// Unlike other languages, Go does not automatically fall through cases.
	switch x {
	case 2:
		fmt.Println("x == 2") // (Matched) Output: x == 2
		fallthrough
	case y + 1:
		fmt.Println("x == y + 1") // (Also executed due to fallthrough) Output: x == y + 1
	}

	// Switch Without an Expression
	// When no expression is provided, the switch acts like a series of "if-else" statements.
	switch {
	case x == 1:
		fmt.Println("x == 1")
	case x == 2:
		fmt.Println("x == 2") // (Matched) Output: x == 2
	default:
		fmt.Println("?")
	}
}
