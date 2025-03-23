// String
// A string is a sequence of characters (runes) enclosed in double quotes.
// Strings are immutable, meaning their values cannot be changed after they are created.
// However, you can create new strings by concatenating existing strings.

package datatypes

import "fmt"

// Creating Strings
// There are some approaches to create strings
func CreatingStrings() {

	// Creating String
	var _ string = "Hello World!"

	// Creating String (Type Inference)
	var _ = "Hello World!"

	// Creating String (Multi-line)
	var _ = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed et aliquam 
		magna, eu dignissim ante. Donec non semper lectus. Vivamus vel efficitur 
		odio. Integer eu pulvinar augue.
	`

	// Creating String (Hexadecimal)
	// Use the "\x<n>" to define hexadecimal characters
	var _ = "\x41" // A

	// Creating String (Octal)
	// Use the "\<n>" to define octal characters
	var _ = "\101" // A
}

// Concatenating Strings
// We can create other strings by joining two or more strings
func ConcatenatingStrings() {

	// Concatenating Strings
	// The "+" operator is used to concatanate two or more strings
	var _ = "Hello, " + "World!"
}

// Escaping Strings
// Some runes must be escaped to be used in a string
func EscapingStrings() {

	// Escaping Strings
	var _ = "\a" // Alert or bell
	var _ = "\"" // Escape Double quote (String)
	var _ = '\'' // Escape Single quote (Rune)
	var _ = "\\" // Backslash
	var _ = "\n" // New line
	var _ = "\t" // Tab
	var _ = "\r" // Carriage return
	var _ = "\b" // Backspace
	var _ = "\f" // Form feed
	var _ = "\v" // Vertical tab

	// Escaping Example (Double Quotes)
	x := "Hello, \"World!\""
	fmt.Println(x) // Hello, "World!"
}

// String Length
// To retrieve the length of the string, the len() func can be used
func StringLength() {

	// String Length (len())
	length := len("Hello, World!")
	fmt.Println(length) // 13
}

// String Indexing
// We can access parts of a string using indexes
func StringIndexing() {

	// Creating String
	// For the examples below, we will need a string
	x := "Hello, World!"

	// String Indexing (Char Code)
	fmt.Println(x[0]) // 72

	// String Indexing (ASCII Value)
	fmt.Println(string(x[0])) // H

	// String Slicing
	// We can define a range using [low:high] to get parts of a string
	fmt.Println(x[:5])  // Hello
	fmt.Println(x[7:])  // World!
	fmt.Println(x[2:4]) // ll
	fmt.Println(x[:])   // Hello, World! (Same reference)
	fmt.Println(x[0:0]) // (Empty string)
}

// String Iteration
// We can use loops to iterate over a string
func StringIteration() {

	// Creating String
	// For the examples below, we will need a string
	x := "Hello, World!"

	// Iterating over a string (for)
	// The example uses a normal for
	for i := 0; i < len(x); i++ {
		fmt.Println(string(x[i])) // Hello, World!
	}

	// Iterating over a string (for range)
	// The example uses the for range
	for _, c := range x {
		fmt.Println(string(c)) // Hello, World!
	}
}
