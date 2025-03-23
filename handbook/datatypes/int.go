// Integer
// An integer (int) is a fundamental numeric datatype in Go, used to
// represent whole numbers.

package datatypes

import "fmt"

// Creating Integers
// We can create integers using different approaches
func CreatingIntegers() {

	// Creating Integer
	// Int is an alias to int32 or int64 depending on the system
	var _ int = 1

	// Creating Integer (Type Inference)
	// Type inference results on an int type by default
	var _ = 1

	// Creating Integer (Binary Notation)
	// Use the "0b" prefix to define integers in binary notation
	var _ = 0b1010 // 10

	// Creating Integer (Hexadecimal Notation)
	// Use the "0x" prefix to define integers in hexdecimal notation
	var _ = 0xFFF // 4095

	// Creating Integer (Octal Notation)
	// Use the "0o" prefix to define integers in octal notation.
	var _ = 0o777 // 511

	// Creating Integer (With Underscore)
	// We can use the underscore "_" to improve readability
	var _ = 100_000_000 // 100000000

	// Creating Integer (With Expression)
	// The result of an expression can be an int
	var _ = 2 + 3*4 // 14
}

// Byte and Rune
// The byte and rune types are aliases for integer types in Go.
// byte is an alias to uint8
// rune is an alias to int32
func ByteAndRune() {

	// Creating Byte (Alias to uint8)
	// Represents a single ASCII character (0-255).
	var b byte = 1

	// Creating Rune (Alias to int32)
	// Represents a Unicode code point.
	var r rune = 1

	// Compare types
	// byte is the same of uint8
	// rune is the same of int32
	fmt.Println(b == uint8(b)) // Output: true
	fmt.Println(r == int32(r)) // Output: true
}

// Operations on Integers
// Integers in Go can be of different sizes, such as int8, int16, int32, int64, and their unsigned counterparts (uint8, uint16, uint32, uint64).
// Here, we will show basic operations with the standard int type.
func OperationsOnIntegers() {

	// Creating Integer Values
	x := 10
	y := 5

	// Addition
	add := x + y
	fmt.Println(add) // Output: 15

	// Subtraction
	sub := x - y
	fmt.Println(sub) // Output: 5

	// Multiplication
	mul := x * y
	fmt.Println(mul) // Output: 50

	// Division
	div := x / y
	fmt.Println(div) // Output: 2

	// Modulo (Remainder)
	mod := x % y
	fmt.Println(mod) // Output: 0
}

// Integer Conversion
// An int can be converted to other int types (int8, byte, rune)
func IntegerConversion() {

	// Widening conversion (expanding the size)
	// This example expands the size from 8-bit to 16-bit
	var x1 int8 = 127
	var y1 int16 = int16(x1)
	fmt.Println(x1, y1) // Output: 127 127

	// Narrowing conversion (truncating excess bits)
	// This example reduces the size from 16-bit to 8-bit
	var x2 int16 = 32767
	var y2 int8 = int8(x2)
	fmt.Println(x2, y2) // Output: 32767 127

	// Signed to unsigned conversion
	// This example changes an signed integer to unsigned
	var x3 int8 = -1
	var y3 uint8 = uint8(x3)
	fmt.Println(x3, y3) // Output: -1 255

	// Unsigned to signed conversion
	// This example changes an unsigned integer to signed
	var x4 uint8 = 255
	var y4 int8 = int8(x4)
	fmt.Println(x4, y4) // Output: 255 -1
}
