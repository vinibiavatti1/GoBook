// Singleton
// The Singleton pattern ensures that a class has only one instance and provides a global point of access to it.
// This is useful when exactly one object is needed to coordinate actions across the system.
// In Go, we can implement the Singleton pattern using a package-level variable and a function to access it.

package gof

import "fmt"

// Declaring a Struct
type Singleton struct {
	ID int
}

// Declaring the Singleton Instance
// Declaring a package-level variable to hold the single instance of the Singleton struct.
var instance *Singleton

// Declaring a Function
// The GetInstance function returns the single instance of the Singleton struct.
func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{ID: 1}
	}
	return instance
}

// Test Singleton
// We can check if the instance is the same by comparing the IDs of the instances returned by GetInstance.
func TestSingleton() {

	// Getting the Singleton instance
	// We will call the GetInstance function twice to get two instances of the Singleton struct.
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	// Check
	// We can see that both instances have the same ID, which means they are indeed the same instance.
	fmt.Println("Singleton Instance IDs:", singleton1.ID, singleton2.ID) // Output: Singleton Instance IDs: 1 1
}
