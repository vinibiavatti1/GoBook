// Type
// The "type" keyword is used to declare new types or create aliases for existing types.
// Syntax:
//   New Type: type <name> <type>
//   Alias:    type <name> = <type>

package syntax

import "fmt"

// Declaring a Type
// We can declare new types using the "type" keyword.
// The new type can be based on an existing type.
// In the example below, LocalType is a new type based on the int type.
// Aliases can be created using the "=" operator.
type LocalType int   // New type
type AliasType = int // Alias (Using "=" operator)

// Declaring Multiple Types
// We can use a type block to declare multiple types at once.
// The examples below shows different base types that can be used for new types.
type (
	Number       float64            // New datatype
	Person       struct{}           // New Struct type
	Handler      interface{}        // New Interface type
	Sum          func(x, y int) int // New Func type
	Point        [2]int             // New Array type
	Collection   []any              // New Slice type
	Stack[T any] []T                // New Slice type (With Generics)
	Dictionary   map[any]any        // New Map type
	Decimal      = float64          // Alias (Using "=" operator)
)

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
