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
	"book/enums"
	"book/syntax"
)

// Main function
// Uncomment the lines below to run the examples in this file.
func main() {

	/*
		Datatypes
	*/

	// Datatypes - Bool
	datatypes.DeclaringBooleans()

	// Datatypes - String
	datatypes.DeclaringStrings()
	datatypes.ManipulatingStrings()

	// Datatypes - Int
	datatypes.DeclaringIntegers()
	datatypes.IntegerOperations()
	datatypes.IntegerConversions()

	// Datatypes - Float
	datatypes.DeclaringFloats()
	datatypes.FloatOperations()
	datatypes.FloatConversion()

	// Datatypes - Complex
	datatypes.DeclaringComplexNumbers()
	datatypes.ManipulatingComplexNumbers()
	datatypes.ComplexNumberOperations()
	datatypes.ComplexNumbersConversion()

	// Datatypes - Any
	datatypes.DeclaringAnyType()

	// Datatypes - Pointer
	datatypes.UsingPointers()

	/*
		Composite
	*/

	// Composite - Array
	composite.DeclaringArrays()
	composite.ManipulatingArrays()

	// Composite - Slice
	composite.DeclaringSlices()
	composite.ManipulatingSlices()

	// Composite - Map
	composite.DeclaringMaps()
	composite.ManipulatingMaps()

	// Composite - Matrix
	composite.DeclaringMatrices()
	composite.ManipulatingMatrices()

	/*
		Syntax
	*/

	// Syntax - Operators
	syntax.ArithmeticOperators()
	syntax.AssignmentOperators()
	syntax.ComparisonOperators()
	syntax.LogicalOperators()
	syntax.AddressOperators()
	syntax.BitwiseOperators()

	// Syntax - Var
	syntax.DeclaringVariables()

	// Syntax - Const
	syntax.DeclaringConstants()

	// Syntax - Iota
	syntax.UsingIotaInFunctions()

	// Syntax - Comments
	syntax.GoDocExample()

	// Syntax - If
	syntax.IfWithComparisonOperators()
	syntax.IfWithLogicalOperators()
	syntax.IfElse()
	syntax.IfWithDeclarations()

	// Syntax - For
	syntax.CreatingLoops()
	syntax.ControllingLoops()
	syntax.IteratingOverData()

	// Syntax - Switch
	syntax.Switches()

	// Syntax - GoTo
	syntax.SimpleGoto()
	syntax.LoopWithGoto()
	syntax.ErrorHandlingWithGoto()
	syntax.BreakingOutLoops()

	// Syntax - Functions
	syntax.LambdaFunctions()
	syntax.UsingClosures()

	// Syntax - Type
	syntax.UsingDeclaredTypes()

	// Syntax - Defer
	syntax.PerformDefer()
	syntax.PerformMultipleDefers()
	syntax.PerformDeferArguments()
	syntax.PerformDeferLoop()
	syntax.PerformDeferFunction()
	syntax.PerformDeferNamedReturn()

	// Syntax - Recursion
	syntax.PerformRecursiveFactorial()

	/*
		Enums
	*/

	// Enums - Enum
	enums.UsingEnum()
}
