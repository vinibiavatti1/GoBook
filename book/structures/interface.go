// Interface
// Interfaces are used to define a set of methods that a type must implement.
// They provide a way to specify the behavior of an object: if something can do this, then it can be used here.
// There is also Type Interface, which is a type that specifies a set of types that a value can have.

package datatypes

import "fmt"

// Interface
// An interface is a type that specifies a set of methods that a concrete type must implement.
type Printer interface {
	Print()
}

// Nil Interface
// Go allows to assign nil to an interface.
// This means that the interface does not hold any value.
var _ Printer = nil

// Defining a Struct
// Methods are implemented on structs.
type Data struct {
	Key   string
	Value string
}

// Implementing Interface
// A type implements an interface by implementing its methods.
// The interface is implicitly implemented by any type that implements all the methods.
func (d *Data) Print() {
	fmt.Printf("%s: %s", d.Key, d.Value)
}

// Using Interface Implementation
// Now we can use the interface to call the method on any type that implements it.
func CheckInterfaceImplementation() {

	// Creating an Instance
	// We will create an instance of the Data struct.
	data := &Data{Key: "A", Value: "1"}

	// Calling the Interface Method
	// Since Data implements the Printer interface, we can call the Print method on it.
	data.Print() // Output: A: 1

	// Checking Type
	// As we can see, we can add Data to a slice of Printer.
	// This is because Data implements the Printer interface.
	slc := []Printer{data}
	fmt.Println(slc) // Output: [&{A 1}]
}

// Local Interfaces
// We can define local interfaces inside functions.
func LocalInterfaces() {

	// Creating an Instance
	// We will create an instance of the Data struct.
	data := &Data{Key: "A", Value: "1"}

	// Defining an Interface
	// We can define an interface inside a function.
	// This interface will only be available inside the function.
	type otherPrinter interface {
		Print()
	}

	// Checking Type
	// Since otherPrinter is the same as Printer, we can assign data to it.
	slc := []otherPrinter{data}
	fmt.Println(slc) // Output: [&{A 1}]
}
