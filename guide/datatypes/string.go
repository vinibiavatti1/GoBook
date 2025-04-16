// String
// A string is a sequence of characters (runes) enclosed in double quotes.
// Strings are immutable, meaning their values cannot be changed after they are created.
// We can manipulate strings to create new strings.

package datatypes

import "fmt"

// Declaring Strings
// There are some approaches to declare strings.
func DeclaringStrings() {

	// Declaring a String
	x := "Hello World!"
	fmt.Println("x:", x) // Output: x: Hello World!

	// Declaring a Multi-line String
	// We can use the backticks (`) to create a multi-line string
	x = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed et aliquam 
		magna, eu dignissim ante. Donec non semper lectus. Vivamus vel efficitur 
		odio. Integer eu pulvinar augue.
	`
	fmt.Println("x:", x) // Output: x: Lorem ipsum dolor sit amet...

	// Declaring a String using Hexadecimal Notation
	// We can use the "\x<n>" syntax to define hexadecimal characters
	x = "\x41"
	fmt.Println("x:", x) // Output: x: A

	// Declaring a String using Octal Notation
	// We can use the "\<n>" syntax to define octal characters
	x = "\101"
	fmt.Println("x:", x) // Output: x: A

	// Concatenating Strings
	// We can create other strings by joining two or more strings
	// The "+" operator is used to concatenate strings
	x = "Hi, " + "John"
	fmt.Println("x:", x) // Output: x: Hi, John

	// Escaping Strings
	// Some runes must be escaped to be used in a string
	x = "\a"                          // Alert or bell
	x = "\""                          // Escape Double quote (String)
	x = "\\"                          // Backslash
	x = "\n"                          // New line
	x = "\t"                          // Tab
	x = "\r"                          // Carriage return
	x = "\b"                          // Backspace
	x = "\f"                          // Form feed
	x = "\v"                          // Vertical tab
	fmt.Println("Escape \"Example\"") // Output: Escape "Example"
}

// Manipulating Strings
// We can use indexing, functions and loops to manipulate strings
func ManipulatingStrings() {

	// Declaring a String
	// For the examples below, we will need a string
	x := "Hello, World!"

	// Retrieving String Length
	// We can use the "len()" function to get the length of a string
	l := len(x)
	fmt.Println("Length:", l) // Output: Length: 13

	// String Indexing
	// We can access a character in a string using indexing
	// Syntax: x[<index>]
	// Note: Since the value is a byte, we need to convert it to a string to get the ASCII representation
	fmt.Println("x[0]:", x[0])         // Output: x[0]: 72
	fmt.Println("x[0]:", string(x[0])) // Output: x[0]: H

	// String Slicing
	// We can access a substring in a string using slicing
	// Note: The end index is exclusive.
	// Syntax: x[<start>:<end>]
	fmt.Println("x[ :5]", x[:5])  // Output: Hello
	fmt.Println("x[7: ]", x[7:])  // Output: World!
	fmt.Println("x[2:4]", x[2:4]) // Output: ll
	fmt.Println("x[ : ]", x[:])   // Output: Hello, World! (Same reference)
	fmt.Println("x[0:0]", x[0:0]) // Output: (empty string)

	// Iterating Strings (Using standard for)
	// We can use the standard for loop to iterate over a string
	for i := 0; i < len(x); i++ {
		c := string(x[i])
		fmt.Println("char:", c) // Output: char: H e l l o...
	}

	// Iterating Strings (Using for-range)
	// The for-range loop is a more concise way to iterate over a string
	for _, c := range x {
		fmt.Println("char:", string(c)) // Output: char: H e l l o...
	}
}
