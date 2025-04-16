// Complex
// A complex number has two parts: a real part and an imaginary part.
// The real and imaginary parts are floating-point numbers.
// Go provides two types for complex numbers:
// - complex64: a complex number where both real and imaginary parts are of type float32.
// - complex128: a complex number where both real and imaginary parts are of type float64 (default type).

package datatypes

import "fmt"

// Declaring Complex Numbers
// We can create complex numbers using the "complex()" builtin function, or directly
// using the "i" notation.
func DeclaringComplexNumbers() {

	// Declaring a Complex Number using "complex()"
	// The "complex()" function takes two float64 arguments: real and imaginary parts.
	// The complex128 is the default type for complex numbers.
	x := complex(1, 2)
	fmt.Println("x:", x) // Output: x: (1+2i)

	// Declaring a Complex Number using "i" Notation
	// The "i" notation is used to define the imaginary part of a complex number.
	x = 1 + 2i
	fmt.Println("x:", x) // Output: x: (1+2i)

	// Declaring a Complex Number of 64-bit Type
	// We can specify the type of a complex number by using the complex64 type.
	var y complex64 = complex(1, 2)
	fmt.Println("y:", y) // Output: y: (1+2i)
}

// Manipulating Complex Numbers
// We can use built-in functions to manipulate complex numbers.
func ManipulatingComplexNumbers() {

	// Declaring a Complex Number
	// We will need a complex number for the examples below.
	x := 1 + 2i

	// Acessing Real Part
	// We can use the "real()" builtin function to access the real part of a complex number.
	r := real(x)
	fmt.Println("Real Part:", r) // Output: Real Part: 1

	// Acessing Imaginary Part
	// We can use the "imag()" builtin function to access the imaginary part of a complex number.
	i := imag(x)
	fmt.Println("Imaginary Part:", i) // Output: Imaginary Part: 2
}

// Operations on Complex Numbers
// We can perform arithmetic operations like addition, subtraction, multiplication, and
// division with complex numbers.
// Modulo operation is not supported for complex numbers.
func ComplexNumberOperations() {

	// Declaring Complex Numbers
	// We will need two complex numbers for the examples below.
	x := 1 + 2i
	y := 3 + 4i

	// Addition
	// The "+" operator is used to add two complex numbers.
	z := x + y
	fmt.Println("z:", z) // Output: z: (4+6i)

	// Subtraction
	// The "-" operator is used to subtract two complex numbers.
	z = x - y
	fmt.Println("z:", z) // Output: z: (-2-2i)

	// Multiplication
	// The "*" operator is used to multiply two complex numbers.
	z = x * y
	fmt.Println("z:", z) // Output: z: (-5+10i)

	// Division
	// The "/" operator is used to divide two complex numbers.
	z = x / y
	fmt.Println("z:", z) // Output: z: (0.44+0.08i)
}

// Complex Numbers Conversion
// Complex numbers can be converted between complex64 and complex128 types using type conversion.
func ComplexNumbersConversion() {

	// Widening conversion (expanding the size)
	// This example expands the size from 64-bit to 128-bit.
	var x1 complex64 = complex(1, 2)
	var y1 complex128 = complex128(x1)
	fmt.Println("complex64:", x1, "complex128:", y1) // Output: complex64: (1+2i) complex128: (1+2i)

	// Narrowing conversion (truncating excess bits)
	// This example reduces the size from 128-bit to 64-bit.
	var x2 complex128 = complex(1, 2)
	var y2 complex64 = complex64(x2)
	fmt.Println("complex128:", x2, "complex64:", y2) // Output: complex128: (3+4i) complex64: (3+4i)
}
