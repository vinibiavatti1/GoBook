// Type
// The "type" keyword is used to declare new types, or to declare aliases for
// existent types.
// Syntax:
//   type <name> <type>

package statements

import "fmt"

// Declaring Structs
// We can use "type" to declare structs.
type Person struct{}

// Declaring Interfaces
// We can use "type" to declare interfaces.
type Handler interface{}

// Declaring Function Signature
// We can use "type" to declare a function signature.
type Sum func(x, y int) int

// Declaring Array Signature
// We can use "type" to declare new array signatures.
type Point [2]int

// Declaring Slice Signature
// We can use "type" to declare new slice signatures.
type Collection []any

// Declaring Map Signature
// We can use "type" to declare new map signatures.
type Dictionary map[any]any

// Declaring Aliases
// We use "type" to declare aliases for existent types.
type Number float64

// Using Declared Types
// This function shows examples of the usage of the types declared above.
func UsingDeclaredTypes() {

	// Using Declared Struct
	person := &Person{}
	fmt.Println(person) // Output: Person{}

	// Using Declared Interface
	var handler Handler = &Person{}
	fmt.Println(handler) // Output: Person{}

	// Using Declared Function
	var sum Sum = func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(7, 3)) // Output: 10

	// Using Declared Array
	p := Point{1, 2}
	fmt.Println(p) // Output: [1, 2]

	// Using Declared Slice
	c := Collection{"A", "B", "C"}
	fmt.Println(c) // Output: [A B C]

	// Using Declared Map
	d := Dictionary{"A": 1}
	fmt.Println(d) // Output: {"A": 1}

	// Using Declared Alias
	var n Number = 3.14
	fmt.Println(n) // Output: 3.14
}
