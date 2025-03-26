// Type
// The "type" keyword is used to declare new types or create aliases for existing types.
// Syntax:
//   type <name> <type>

package statements

import "fmt"

// Declaring Structs
// We use "type" to declare structs.
type Person struct{}

// Declaring Interfaces
// Interfaces define a set of methods that a type must implement.
type Handler interface{}

// Declaring Function Signature
// "type" can be used to define function signatures.
type Sum func(x, y int) int

// Declaring Array Signature
// "type" can be used to create new names for fixed-size arrays.
type Point [2]int

// Declaring Slice Signature
// We can create custom names for slices.
type Collection []any

// Declaring Map Signature
// We create custom names for maps.
type Dictionary map[any]any

// Declaring Aliases
// We create an alias for an existing type.
type Number float64

// Using Declared Types
// This function shows examples of the usage of the types declared above.
func UsingDeclaredTypes() {

	// Using Declared Struct
	person := &Person{}
	fmt.Println(person) // Output: &{}

	// Using Declared Interface
	var handler Handler = &Person{}
	fmt.Println(handler) // Output: &{}

	// Using Declared Function
	var sum Sum = func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(7, 3)) // Output: 10

	// Using Declared Array
	p := Point{1, 2}
	fmt.Println(p) // Output: [1 2]

	// Using Declared Slice
	c := Collection{"A", "B", "C"}
	fmt.Println(c) // Output: [A B C]

	// Using Declared Map
	d := Dictionary{"A": 1}
	fmt.Println(d) // Output: map[A:1]

	// Using Declared Alias
	var n Number = 3.14
	fmt.Println(n) // Output: 3.14
}
