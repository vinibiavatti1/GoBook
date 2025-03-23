// Constants
// In Go, constants are values that cannot be changed once set. Constants can be of any type,
// including numeric, string, or boolean types.
// They are declared using the const keyword, and their values are determined at compile time
// Constants are typically used for values that are not expected to change during the execution.
// of the program
// Note: The value of a constant must be determinable at compile time.

package statements

// Global Constants
// These constants are globals and can be accessed by any function
// inside the package.
const _ int = 1
const _ = 1

// Multiple Global Constants
// We can use a const block to define multiple constants.
const (
	Val = 1
	_   = 1 // Unused
)

// Multiple Global Constants With Same Value
// Constants without a literal value specified inherit the same value of
// above constant.
const (
	Val1 = 1 // 1
	Val2     // 1
	Val3 = 2 // 2
	Val4     // 2
)

// Declaring Constants
// There are several ways to declare constants.
func DeclaringConstants() {

	// Declaring Constants
	const x1 int = 10
	println(x1) // Output: 10

	// Constants Are Immutable
	// x1 = 20 = Compile Error

	// Declaring Constants (Type Inference)
	const x2 = 10
	println(x2) // Output: 10

	// Declaring Constants (:=)
	x3 := 10
	println(x3) // Output: 10

	// Declaring Inline Constants (Inline)
	const a1, b1, c1 int = 1, 2, 3
	println(a1, b1, c1) // Output: 1 2 3

	// Declaring Inline Constants (Type Inference)
	const a2, b2, c2 = 1, true, 3.5
	println(a2, b2, c2) // Output: 1 true 3.5

	// Declaring Inline Constants (:=)
	a3, b3, c3 := 1, true, 3.5
	println(a3, b3, c3) // Output: 1 true 3.5

	// Declaring Multiple Constants
	const (
		a4 = 1
		b4 = true
		c4 = "Hello"
	)
	println(a4, b4, c4) // Output: 1 true Hello

	// Declaring Unused Constants
	const _ int = 1

	// Declaring Unused Constants (Type Inference)
	const _ = 1
}
