// Type
// The "type" keyword is used to declare new types or create aliases for existing types.
// Syntax:
//   type <name> <type>

package syntax

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

// Declaring New Type Based on Other Type
// We can create a new type based on an existing type.
// In this case, Number is not a float64, but a new type that is based on it.
type Number float64

// Declaring Alias
// We can create an alias for an existing type.
// Note: The "=" operator is used to define an alias.
// In this case, Decimal is a float64 (same type).
type Decimal = float64

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

	// Using New Type
	var n Number = 3.14
	fmt.Println(n) // Output: 3.14

	// Using Alias
	var x Decimal = 3.14
	fmt.Println(x) // Output: 3.14
}

// Local Types
// Local types are types that are declared within a function scope.
func LocalTypes() {

	// Declaring Local Type
	// We can declare a type within a function.
	// This type is not accessible outside the function.
	type value int

	// Using Local Type
	// We can use the local type within the function.
	var v value = 42
	fmt.Println(v) // Output: 42
}
