// Float
// The float type in Go is used to represent floating-point numbers. It is used to store
// decimal values with variable precision
// Go provides two primary types for floating-point numbers: float32 and float64
// The float32 has 32-bit precision and the float64 has 64-bit precision, with float64
// being the default for floating-point numbers

package datatypes

import "fmt"

// Creating Floats
// We can create floating-point numbers using different approaches
func CreatingFloats() {

	// Creating Float (float64 is the default type)
	// By default, Go assigns float64 type to floating-point numbers
	var _ float64 = 3.14159

	// Creating Float (Explicit Type Declaration)
	// We can also explicitly declare a float type
	var _ float32 = 3.14

	// Creating Float (Type Inference)
	// Type inference works as well for float types
	var _ = 2.71828

	// Creating Float (Scientific Notation)
	// Go allows scientific notation using "e" or "E"
	var _ = 1.23e4 // 12300
}

// Float Precision
// Floating-point numbers are not always perfectly accurate due to the limitations of
// binary floating-point representation.
func FloatPrecision() {

	// Floating-point precision issues
	// Example of precision loss in float64
	var x float64 = 0.1 + 0.2
	var y float64 = 0.3
	fmt.Println(x == y) // Output: false (due to floating-point precision)
}

// Operations on Floats
// Floats in Go are represented by the `float32` and `float64` types, with `float64` being the default.
// Here, we will show basic operations with the standard `float64` type.
func OperationsOnFloats() {

	// Creating Float Values
	x := 10.5
	y := 5.2

	// Addition
	add := x + y
	fmt.Println(add) // Output: 15.7

	// Subtraction
	sub := x - y
	fmt.Println(sub) // Output: 5.3

	// Multiplication
	mul := x * y
	fmt.Println(mul) // Output: 54.6

	// Division
	div := x / y
	fmt.Println(div) // Output: 2.019230769230769
}

// Converting Floats
// Floats can be converted to other numeric types (e.g., int, float32, float64)
func FloatConversion() {

	// Convert float64 to int (truncates the fractional part)
	var x float64 = 3.99
	var y int = int(x)
	fmt.Println(x, y) // Output: 3.99 3

	// Convert float64 to int (for negative numbers)
	var x2 float64 = -3.99
	var y2 int = int(x2)
	fmt.Println(x2, y2) // Output: -3.99 -3

	// Convert float32 to float64
	var x3 float32 = 3.14
	var y3 float64 = float64(x3)
	fmt.Println(x3, y3) // Output: 3.14 3.14

	// Convert float64 to float32 (may lose precision)
	var x4 float64 = 3.141592653589793
	var y4 float32 = float32(x4)
	fmt.Println(x4, y4) // Output: 3.141592653589793 3.141593
}
