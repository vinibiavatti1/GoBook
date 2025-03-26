// Book
// GoBook is a comprehensive guide to learning the Go programming language.
// This repository is packed with detailed examples, explanations, and best practices
// to help developers of all levels master Go. From basic syntax and data types to
// advanced topics like concurrency and interfaces, GoBook covers it all with
// practical, real-world examples.

package main

import (
	"book/composite"
	"book/datatypes"
	"book/expressions"
	"book/statements"
	"fmt"
)

// Main function
func main() {
	// datatypes/bool.go
	fmt.Println("> datatypes/bool.go")
	datatypes.DeclaringBooleans()

	// datatypes/string.go
	fmt.Println("> datatypes/string.go")
	datatypes.DeclaringStrings()
	datatypes.ManipulatingStrings()

	// datatypes/int.go
	fmt.Println("> datatypes/int.go")
	datatypes.DeclaringIntegers()
	datatypes.IntegerOperations()
	datatypes.IntegerConversions()

	// datatypes/float.go
	fmt.Println("> datatypes/float.go")
	datatypes.DeclaringFloats()
	datatypes.FloatOperations()
	datatypes.FloatConversion()

	// datatypes/complex.go
	fmt.Println("> datatypes/complex.go")
	datatypes.DeclaringComplexNumbers()
	datatypes.ManipulatingComplexNumbers()
	datatypes.ComplexNumberOperations()
	datatypes.ComplexNumbersConversion()

	// composite/array.go
	fmt.Println("> composite/array.go")
	composite.DeclaringArrays()
	composite.ManipulatingArrays()

	// composite/slice.go
	fmt.Println("> composite/slice.go")
	composite.DeclaringSlices()
	composite.ManipulatingSlices()

	// composite/map.go
	fmt.Println("> composite/map.go")
	composite.DeclaringMaps()
	composite.ManipulatingMaps()

	// expressions/operators.go
	fmt.Println("> expressions/operators.go")
	expressions.ArithmeticOperators()
	expressions.AssignmentOperators()
	expressions.ComparisonOperators()
	expressions.LogicalOperators()
	expressions.AddressOperators()
	expressions.BitwiseOperators()

	// statements/var.go
	fmt.Println("> statements/var.go")
	statements.DeclaringVariables()

	// statements/const.go
	fmt.Println("> statements/const.go")
	statements.DeclaringConstants()

	// statements/comments.go
	fmt.Println("> statements/comments.go")
	statements.GoDocExample()
}
