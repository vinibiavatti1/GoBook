// Complex Number
// A complex number has two parts: a real part and an imaginary part.
// The real and imaginary parts are floating-point numbers.
//
// Go provides two types for complex numbers:
// - complex64: a complex number where both real and imaginary parts are of type float32
// - complex128: a complex number where both real and imaginary parts are of type float64

package datatypes

import "fmt"

// Creating Complex Numbers
// We can create complex numbers using the complex function or directly with the i notation.
func CreatingComplexNumbers() {

	// Creating Complex Number (complex128)
	// Real: 3.0, Imaginary: 4.0
	var _ complex128 = complex(3.0, 4.0)

	// Creating Complex Number (complex64)
	// Real: 1.5, Imaginary: 2.5
	var _ complex64 = complex(1.5, 2.5)

	// Creating Complex Number (i Notation)
	// Real: 3, Imaginary: 4
	var _ complex128 = 3 + 4i

	// Creating Complex Number (Type Inference)
	// Type inference results on a complex128 type by default
	var _ = 3 + 4i
}

// Accessing Real and Imaginary Parts
// We can access the real and imaginary parts of a complex number using the real() and imag() functions.
func AccessingRealAndImaginaryParts() {

	// Creating Complex Number
	x := 3 + 4i

	// Accessing Real Part (real())
	var _ = real(x) // 3

	// Acessing Imaginary Part (imag())
	var _ = imag(x) // 4
}

// Operations on Complex Numbers
// We can perform arithmetic operations like addition, subtraction, multiplication, and
// division with complex numbers.
func OperationsOnComplexNumbers() {

	// Creating Complex Numbers
	x := 3 + 4i
	y := 1 + 2i

	// Addition
	add := x + y
	fmt.Println(add) // Output: (4 + 6i)

	// Subtraction
	sub := x - y
	fmt.Println(sub) // Output: (2 + 2i)

	// Multiplication
	mul := x * y
	fmt.Println(mul) // Output: (-5 + 10i)

	// Division
	div := x / y
	fmt.Println(div) // Output: (2 + 0i)
}

// Converting Between Complex Types
// Complex numbers can be converted between complex64 and complex128 types by type conversion.
func ComplexTypeConversion() {

	// Converting complex128 to complex64
	var x1 complex128 = complex(3.0, 4.0)
	var y1 complex64 = complex64(x1)
	fmt.Println(x1, y1) // (3 + 4i) (3 + 4i)

	// Converting complex64 to complex128
	var x2 complex64 = complex(1.5, 2.5)
	var y2 complex128 = complex128(x2)
	fmt.Println(x2, y2) // (1.5 + 2.5i) (1.5 + 2.5i)
}
