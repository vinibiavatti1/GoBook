// Styleguide
// This file will explain the convention of Go language, the rules of writing Go code, and the best practices.
// The Go programming language is designed to be simple and efficient, and its conventions reflect that philosophy.
// The conventions are not just about syntax, but also about how to structure your code, how to name your variables
// and functions, and how to organize your packages.
// Go Style Guide: https://google.github.io/styleguide/go

package styleguide

import "fmt"

/*
	Naming Convention
*/

// Public Modules
// Characters: a-z, 0-9, /, .
// Pattern: "github.com/<username>/<repo>"
var _ = "github.com/john/myproject"

// Private Modules
// Characters: a-z, 0-9, /, .
var _ = "myproject"

// Directories
// Characters: a-z, 0-9
var _ = "httphandler"

// Files
// Characters: a-z, 0-9
// Extension: .go
var _ = "httphandler.go"

// Test Files
// Characters: a-z, 0-9
// Extension: .go
// Pattern: "<name>_test.go"
var _ = "httphandler_test.go"

// Packages
// Characters: a-z, 0-9
// Note: The package name should be the same as the directory name where the package is located.
var _ = "httphandler"

// Acronyms
// Acronyms should be written in all uppercase letters.
// Characters: A-Z, 0-9
var ID = 1
var HTTPClient = 1

// Keywords
// When using keywords as resource names, they should be suffixed with an underscore "_".
func If_() {}

// Unused
// Unused resources should have the "_" name.
var _ = 1

// Variables
// Characters: a-z, A-Z, 0-9
var PersonID = 1

// Constants
// Characters: a-z, A-Z, 0-9
const EmployeeID = 1

// Functions
// Characters: a-z, A-Z, 0-9
func PerformOperation() {}

// Structs
// Characters: a-z, A-Z, 0-9
type Person struct{ name string }

// Methods
// Characters: a-z, A-Z, 0-9
func (p *Person) JobName() {}

// Type Alias
// Characters: a-z, A-Z, 0-9
type Number int

// Interfaces
// Characters: a-z, A-Z, 0-9
type Reader interface{ Read() }

// Special Names
// The function below shows special variable names that are used in Go language.
// Special Conventions:
// - Error: err
// - Index: i, j, k
// - Key: k
// - Value: v
// - Channel: ch
func SpecialVariableNames() {

	// Error Name
	// The error name should be "err".
	_, err := map[int]int{}[0]
	fmt.Println(err)

	// Index Name
	// The index name should be "i", "j", "k".
	for i := range 3 {
		for j := range 3 {
			fmt.Println(i, j)
		}
	}

	// Value Name
	// The value name should be "v".
	for v := range []int{1, 2, 3} {
		fmt.Println(v)
	}

	// Key-Value Pair
	// The key name should be "k".
	// The value name should be "v".
	for k, v := range map[int]int{1: 2} {
		fmt.Println(k, v)
	}

	// Channel Name
	// The channel name should be "ch".
	ch := make(chan int)
	fmt.Println(ch)
}

/*
	Comment Convention
*/

// Single Line Comment
// Single line comments should always be used to document the resources below.
// - Packages
// - Functions
// - Types (Structs, Interfaces, etc.)
// - Global Variables
// - Global Constants

// Section Comment
// Section comments can be used to separate different sections of code.
// A multi-line must be used to create section comments.
/*
	Section Comment
*/

/*
	Methods Convention
*/

// Method Receiver
// The method receiver should be a pointer to the struct type.
// The receiver name should be a single letter, preferably the first letter of the struct name.
// If the struct name is various words, the receiver name should be the first letter of each word.
func (p *Person) Surname() {}

// Getters
// Functions that return something are given noun-like names.
// Don't use "Get" prefix.
func (p *Person) Name() string { return p.name }

// Setters
// Functions that do something are given verb-like names.
// For setters, we use "Set" prefix.
func (p *Person) SetName(name string) { p.name = name }
