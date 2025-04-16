// Integer
// An integer (int) is a fundamental numeric datatype in Go, used to represent integer numbers.
// Signed integers can represent positive and negative numbers.
// Unsigned integers can only represent positive numbers.
// The byte and rune types are aliases for integer types in Go.

package datatypes

import "fmt"

// Declaring Integers
// We can create integers using different approaches.
func DeclaringIntegers() {

	// Declaring an Integer
	// The "int" type is used to represent integers.
	// The size depends on the platform (32 or 64 bits).
	x := 1
	fmt.Println("x:", x) // Output: x: 1

	// Declaring an Integer Using Binary Notation
	// We can use the "0b<n>" notation to define integers in binary.
	x = 0b1010           // Same as x = 10
	fmt.Println("x:", x) // Output: x: 10

	// Declaring an Integer Using Hexadecimal Notation
	// We can use the "0x<n>" notation to define integers in hexdecimal.
	x = 0xFFF            // Same as x = 4095
	fmt.Println("x:", x) // Output: x: 4095

	// Declaring an Integer Using Octal Notation
	// We can use the "0o<n>" notation to define integers in octal.
	x = 0o777            // Same as x = 511
	fmt.Println("x:", x) // Output: x: 511

	// Declaring an Integer Using a Rune
	// Since a rune is an alias for int32, we can use it to represent integers.
	// A rune is a Unicode code point.
	x = 'A'              // Same as x = 65
	fmt.Println("x:", x) // Output: x: 65

	// Declaring an Integer Using Underscores
	// We can use the underscore "_" to separate digits and improve readability.
	x = 100_000_000      // Same as x = 100000000
	fmt.Println("x:", x) // Output: x: 100000000

	// Declaring an Computed Integer
	// We can create integers from expressions which will be computed to generate the result.
	x = 2 + 3*4
	fmt.Println("x:", x) // Output: x: 14

	// Declaring other Integer Types
	// Other integer types can be used to represent different sizes.
	// The unsigned integer types can only represent positive numbers.
	var (
		_ int8    = 1 // 8-bit: -128 to 127
		_ int16   = 1 // 16-bit: -32768 to 32767
		_ int32   = 1 // 32-bit: -2147483648 to 2147483647
		_ int64   = 1 // 64-bit: -9223372036854775808 to 9223372036854775807
		_ uint    = 1 // 32/64-bit Unsigned integer (default)
		_ uint8   = 1 // 8-bit: 0 to 255
		_ uint16  = 1 // 16-bit: 0 to 65535
		_ uint32  = 1 // 32-bit: 0 to 4294967295
		_ uint64  = 1 // 64-bit: 0 to 18446744073709551615
		_ uintptr = 1 // Memory Address Type (32 or 64 bits)
		_ byte    = 1 // Alias for uint8
		_ rune    = 1 // Alias for int32 (Unicode code point)
	)
}

// Integer Operations
// We will show basic operations with the standard int type like addition, subtraction,
// multiplication, division and modulo.
func IntegerOperations() {

	// Addition Operation
	// The "+" operator is used to add two integers.
	x := 7 + 3
	fmt.Println("x:", x) // Output: x: 10

	// Subtraction Operation
	// The "-" operator is used to subtract two integers.
	x = 7 - 3
	fmt.Println("x:", x) // Output: x: 4

	// Multiplication Operation
	// The "*" operator is used to multiply two integers.
	x = 7 * 3
	fmt.Println("x:", x) // Output: x: 21

	// Division Operation
	// The "/" operator is used to divide two integers.
	x = 9 / 3
	fmt.Println("x:", x) // Output: x: 3

	// Modulo Operation (Remainder)
	// The "%" operator is used to get the remainder of a division.
	x = 3 % 2
	fmt.Println("x:", x) // Output: x: 1
}

// Integer Conversions
// Integers can be converted to other integer types using explicit conversions.
func IntegerConversions() {

	// Widening conversion (expanding the size)
	// This example expands the size from 8-bit to 16-bit.
	var x1 int8 = 127
	var y1 int16 = int16(x1)
	fmt.Println("int8:", x1, "int16:", y1) // Output: int8: 127 int16: 127

	// Narrowing conversion (truncating excess bits)
	// This example reduces the size from 16-bit to 8-bit.
	var x2 int16 = 32767
	var y2 int8 = int8(x2)
	fmt.Println("int16:", x2, "int8:", y2) // Output: int16: 32767 int8: -1

	// Signed to unsigned conversion
	// This example changes an signed integer to unsigned.
	var x3 int8 = -1
	var y3 uint8 = uint8(x3)
	fmt.Println("int8:", x3, "uint8:", y3) // Output: int8: -1 uint8: 255

	// Unsigned to signed conversion
	// This example changes an unsigned integer to signed.
	var x4 uint8 = 255
	var y4 int8 = int8(x4)
	fmt.Println("uint8:", x4, "int8:", y4) // Output: uint8: 255 int8: -1
}
