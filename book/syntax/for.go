// For
// The "for" keyword is used to create loops in Go. It is versatile and can be used in various ways
// to iterate over data or repeat a block of code based on specific conditions.
// Syntax:
//   for {...}
//   for <condition> {...}
//   for <variables> := range <data> {...}
//   for <variables>; <condition>; <increment> {...}

package syntax

import "fmt"

// Creating Loops
// This function demonstrates how to create loops using the "for" keyword.
func CreatingLoops() {

	// For With Condition (While Loop)
	// Go does not have a "while" keyword. Instead, a "for" loop can be used with a condition,
	// which functions similarly to a "while" loop.
	i := 0
	for i < 3 {
		fmt.Println(i) // Output: 0 1 2
		i++
	}

	// For with Index (Standard For)
	// A traditional for loop that repeats a block of code based on a condition.
	for i := 0; i < 3; i++ {
		fmt.Println(i) // Output: 0 1 2
	}

	// For with Multiple Indexes
	// We can use multiple index variables in the loop, which is useful when iterating with two conditions.
	for i, j := 0, 0; i < 3 && j < 3; i, j = i+1, j+1 {
		fmt.Println(i, j) // Output: 0 0 1 1 2 2
	}

	// For with Range (Index Only)
	// We can use the "range" keyword to iterate over various data structures like arrays, slices, strings, and maps.
	// Range returns the index and value of the data structure.
	// In this example, we only use the index.
	x := []string{"A", "B", "C"}
	for i := range x {
		fmt.Println(i) // Output: 0 1 2
	}

	// For with Range (Index and Value)
	// The second variable in the "range" expression is used to get the value of the data structure.
	// In this example, we use both the index and value.
	for i, v := range x {
		fmt.Println(i, v) // Output: 0 A 1 B 2 C
	}

	// For with Range (Value Only)
	// If we don't need the index, we can use the blank identifier "_" to discard it.
	for _, v := range x {
		fmt.Println(v) // Output: A B C
	}

	// For with Range (No Variable)
	// We can also create a for-range without any variable
	for range 3 {
		fmt.Println("A") // Output: A A A
	}

	// Infinite For Loop
	// An empty "for" loop is used to create an infinite loop, which can be controlled using "break".
	// It is the equivalent of a "while true" loop in other languages.
	for {
		break
	}

	// Empty Statements
	// We can use empty declarations, conditions and increments in a "for" loop.
	// The example below demonstrates an empty declaration.
	for ; i > 0; i++ {
		break
	}
}

// Controlling Loops
// This function demonstrates how to control loops using "continue" and "break".
func ControllingLoops() {

	// Continue
	// The "continue" keyword is used to skip the current iteration of the loop and jump to the next iteration.
	// In this example, we skip the iteration when i is equal to 0.
	for i := range 3 {
		if i == 0 {
			continue
		}
		fmt.Println(i) // Output: 1 2
	}

	// Break
	// The "break" keyword is used to exit the loop entirely, even if the condition hasn't been fully met.
	// In this example, we break the loop when i is equal to 1.
	for i := range 3 {
		if i == 1 {
			break
		}
		fmt.Println(i) // Output: 0
	}
}

// Iterating Over Data
// This function demonstrates how to iterate over data using "for" loops.
func IteratingOverData() {

	// Iterating Strings
	// Strings in Go are UTF-8 encoded, so iterating over a string will return the Unicode code points.
	str := "ABC"
	for i, v := range str {
		fmt.Println(i, string(v)) // Output: 0 A 1 B 2 C
	}

	// Iterating Arrays
	// Arrays in Go are fixed-size, so we can iterate over them using a "for" loop.
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v) // Output: 0 1 1 2 2 3
	}

	// Iterating Slices
	// Slices are more flexible than arrays, and we can iterate over them using a "for" loop.
	slc := []int{4, 5, 6}
	for i, v := range slc {
		fmt.Println(i, v) // Output: 0 4 1 5 2 6
	}

	// Iterating Maps
	// Maps in Go are key-value pairs, and we can iterate over them using a "for" loop.
	mp := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for k, v := range mp {
		fmt.Println(k, v) // Output: A 1 B 2 C 3
	}
}
