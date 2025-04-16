// Float
// The float type in Go is used to represent floating-point numbers.
// It is used to store decimal values with variable precision.
// Go provides two primary types for floating-point numbers:
// - float32: a 32-bit floating-point number
// - float64: a 64-bit floating-point number (default type)

package datatypes

import "fmt"

// Declaring Floats
// We can declare floating-point numbers using floating-point numbers.
func DeclaringFloats() {

	// Declaring a Float
	// We can declare a float using a floating-point number.
	// The default type is float64.
	x := 3.14
	fmt.Println("x:", x) // Output: x: 3.14

	// Creating Float with Scientific Notation
	// The scientific notation ("e" or "E") can be used to represent floats.
	// Note that the result is a float64.
	x = 1.23e4
	fmt.Println("x:", x) // Output: x: 12300

	// Creating a 32-bit Float
	// We can explicitly declare a float32 type.
	var y float32 = 3.14
	fmt.Println("y:", y) // Output: y: 3.14

	// Float Precision
	// Floating-point numbers are not always perfectly accurate due to the limitations of
	// binary floating-point representation.
	// Note that in the example below, the result is not exactly 0.3.
	x = 0.1 + 0.2
	fmt.Println("x:", x) // Output: x: 0.30000000000000004
}

// Float Operations
// We can perform arithmetic operations like addition, subtraction, multiplication, and
// division with floating-point numbers.
// Modulo operation is not supported for floating-point numbers.
func FloatOperations() {

	// Addition
	// We can use the "+" operator to add two floats.
	x := 3.0 + 2.5
	fmt.Println("x:", x) // Output: x: 5.5

	// Subtraction
	// We can use the "-" operator to subtract two floats.
	x = 3.0 - 2.5
	fmt.Println("x:", x) // Output: x: 0.5

	// Multiplication
	// We can use the "*" operator to multiply two floats.
	x = 3.0 * 2.5
	fmt.Println("x:", x) // Output: x: 7.5

	// Division
	// We can use the "/" operator to divide two floats.
	x = 5.5 / 2.0
	fmt.Println("x:", x) // Output: x: 2.75
}

// Float Conversions
// Floats can be converted to other numeric types (e.g., int, float32, float64)
func FloatConversion() {

	// Widening conversion (expanding the size)
	// This example expands the size from 32-bit to 64-bit.
	var x1 float32 = 1.5
	var y1 float64 = float64(x1)
	fmt.Println("float32:", x1, "float64:", y1) // Output: float32: 1.5 float64: 1.5

	// Narrowing conversion (truncating excess bits)
	// This example reduces the size from 64-bit to 32-bit.
	var x2 float64 = 3.14
	var y2 float32 = float32(x2)
	fmt.Println("float64:", x2, "float32:", y2) // Output: float64: 3.14 float32: 3.14

	// Converting float to int
	// Floats can be converted to integers by truncating the decimal part.
	var x3 float64 = 3.99
	var y3 int = int(x3)
	fmt.Println("float64:", x3, "int:", y3) // Output: float64: 3.99 int: 3

	// Converting int to float
	// Integers can be converted to floats without losing precision.
	var x4 int = 5
	var y4 float64 = float64(x4)
	fmt.Println("int:", x4, "float64:", y4) // Output: int: 5 float64: 5
}
