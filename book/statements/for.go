// For
// The "for" keyword is used to create loops in Go. It is versatile and can be used in various ways
// to iterate over data or repeat a block of code based on specific conditions.
// Syntax:
//   for <declaration>; <condition>; <increment> {...}
//   for <variables> := range <sequence> {...}

package statements

import "fmt"

// For With Condition (While)
// Go does not have a "while" keyword. Instead, a "for" loop can be used with a condition,
// which functions similarly to a "while" loop.
func ForWithCondition() {

	// For With Condition (acts like a while loop)
	// The loop continues to run as long as the condition (i < 3) is true.
	i := 0
	for i < 3 {
		fmt.Println(i) // Output: 0 1 2
		i++
	}
}

// For With Index
// This is the most common form of a "for" loop, where we specify an index variable.
func ForWithIndex() {

	// Standard For Loop with Index
	// A traditional for loop with an index variable (i) that increments each iteration.
	for i := 0; i < 3; i++ {
		fmt.Println(i) // Output: 0 1 2
	}

	// For Loop with Multiple Indexes
	// You can use multiple index variables in the loop, which is useful when iterating with two conditions.
	for i, j := 0, 0; i < 3 && j < 3; i, j = i+1, j+1 {
		fmt.Println(i, j) // Output: 0 0 1 1 2 2
	}
}

// For With Range
// The "range" keyword allows you to iterate over various data structures like arrays, slices, strings, and maps.
func ForWithRange() {

	// Iterating Over a Range
	// The range expression provides the index values for iterating over slices, arrays, or strings.
	for i := range 3 {
		fmt.Println(i) // Output: 0 1 2
	}

	// Iterating Over a String
	// The "range" keyword iterates over a string, returning the index and the Unicode code point (rune).
	str := "ABC"
	for i, v := range str {
		fmt.Println(i, string(v)) // Output: 0 A 1 B 2 C
	}

	// Iterating Over an Array (or Slice)
	// You can iterate over arrays or slices using "range" to get the index and value.
	nums := []int{9, 8, 7}
	for i, v := range nums {
		fmt.Println(i, v) // Output: 0 9 1 8 2 7
	}

	// Iterating Over a Map
	// When iterating over a map, the loop gives you the key and value of each entry.
	mp := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for k, v := range mp {
		fmt.Println(k, v) // Output: A 1 B 2 C 3
	}

	// Discarding a Variable in Range
	// You can discard a variable if it is not needed in the loop.
	str2 := "ABC"
	for _, v := range str2 {
		fmt.Println(string(v)) // Output: A B C (ignores the index variable)
	}
}

// For Infinite
// A "for" loop with no condition is used to create an infinite loop.
func ForInfinite() {

	// Infinite For Loop
	// An empty "for" loop is used to create an infinite loop, which can be controlled using "break".
	for {
		break
	}
}

// Continue and Break
// The "continue" and "break" keywords are used to control the flow of a loop.
// "continue" skips the current iteration, while "break" exits the loop entirely.
func ContinueAndBreak() {

	// Continue
	// The "continue" keyword is used to skip the current iteration of the loop and jump to the next iteration.
	for i := range 3 {
		if i == 0 {
			continue // Skips the current iteration when i == 0
		}
		fmt.Println(i) // Output: 1 2
	}

	// Break
	// The "break" keyword is used to exit the loop entirely, even if the condition hasn't been fully met.
	for i := range 3 {
		if i == 1 {
			break // Breaks the loop when i == 1
		}
		fmt.Println(i) // Output: 0
	}
}
