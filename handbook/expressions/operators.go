// Operators
// Go provides a rich set of operators to work with different types of data:
// 1. Arithmetic Operators: +, -, *, /, %
// 2. Assignment Operators: =, +=, -=, *=, /=, %=
// 3. Comparison Operators: ==, !=, >, <, >=, <=
// 4. Logical Operators: &&, ||, !
// 5. Address Operators: & *
// 6. Bitwise Operators: &, |, ^, <<, >>

package expressions

import "fmt"

// Arithmetic Operators
// These operators perform arithmetic operations such as addition, subtraction,
// multiplication, and division.
func ArithmeticOperators() {

	// Creating variables
	x, y := 10, 5

	// Arithmetic Operators
	_ = x + y // Addition (15)
	_ = x - y // Subtraction (5)
	_ = x * y // Multiplication (50)
	_ = x / y // Division (2)
	_ = x % y // Modulo (Remainder) (0)
}

// Assignment Operators
// These operators are used to assign values to variables.
func AssignmentOperators() {

	// Creating variables
	x := 10

	// Assignment Operators
	x = 20 // Assignment (20)
	x += 5 // Add and Assign (25)
	x -= 3 // Subtract and Assign (22)
	x *= 2 // Multiply and assign (44)
	x /= 4 // Divide and assign (11)
	x %= 3 // Modulo and assign (2)
}

// Comparison Operators
// These operators compare values and return a boolean result (true or false).
func ComparisonOperators() {

	// Creating variables
	x, y := 10, 5

	// Equal to
	_ = x == y // Equal (false)
	_ = x != y // Not Equal (true)
	_ = x > y  // Greater Than (true)
	_ = x < y  // Less Than (false)
	_ = x >= y // Greater Than or Equal To (true)
	_ = x <= y // Less Than or Equal To (false)
}

// Logical Operators
// These operators perform logical operations, typically used with boolean values.
func LogicalOperators() {

	// Creating boolean variables
	x, y := true, false

	_ = x && y // AND (false)
	_ = x || y // OR (true)
	_ = !x     // NOT (false)
}

// Address Operators
// These operators are used to get the memory address of a variable
// (using &) and to dereference pointers (using *).
func AddressOperators() {

	// Creating a variable
	x := 10

	// Getting the memory address of 'a' (address operator)
	ptr := &x
	fmt.Println(ptr) // Output: 0xc000067f30

	// Dereferencing the pointer (dereference operator)
	fmt.Println(*ptr) // Output: 10
}

// Bitwise Operators
// These operators perform operations at the bit level on integer values.
func BitwiseOperators() {

	// Creating variables
	x := 60 // (0011 1100) in binary
	y := 13 // (0000 1101) in binary

	// AND
	_ = x & y  // AND (0000 1100)
	_ = x | y  // OR (0011 1101)
	_ = x ^ y  // XOR (0011 0001)
	_ = x << 2 // Left Shift (1111 0000)
	_ = x >> 2 // Right Shift (0000 1111)
}
