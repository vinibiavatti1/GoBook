// Error
// In Go, there isn't a try-catch mechanism, instead, errors are handled by checking the error value returned from a
// function.
// Errors in Go are values that implement the error interface.
// The error value is nil if there is no error.

package errors

import (
	"errors"
	"fmt"
)

// Declaring Function that Raises an Error
// The function below shows how to declare a function that returns an error.
// The error can be nil if there is no error.
// We can use the "fmt.Errorf" function to create an error.
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

// Declaring Function that Raises an Error (With Errors.New)
// We can also use the "errors.New" function to create an error.
func Divide2(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

// Handling Error
// We will call the function above to raise an error.
func HandlingError() {

	// Handling Errors
	// We can check if the error value is nil or not.
	// If the error value is not nil, we can handle the error.
	// This is the same approach to validate errors created using "errors.New"
	r, err := Divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: Division by zero
	} else {
		fmt.Println("Result:", r)
	}
}

// Creating a Custom Error
// Any struct that implements the error interface can be used as a custom error.
// The "error" interface is located in the "errors" package.
type CustomError struct {
	Code    int
	Message string
}

// Implementing the Error Interface
// The "Error()" method is used to implement the error interface.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Declaring Function that can Raises a Custom Error
// The function below raises the custom error.
func Request(path string) (string, error) {
	if path == "/admin" {
		return "", &CustomError{
			Code:    403,
			Message: "Forbidden",
		}
	}
	return "Success", nil
}

// Handling a Custom Error
// The function below shows how to handle the custom error.
// We can check the type of the error using type assertion to ensure that the error is of type "CustomError".
func HandlingCustomError() {

	// Handling Custom Error
	// We will call the function above to raise the custom error.
	_, err := Request("/admin")
	if err != nil {
		if _, ok := err.(*CustomError); ok {
			fmt.Println("Error:", err) // Output: Error: Code: 403, Message: Forbidden
		} else {
			fmt.Println("Generic Error")
		}
	} else {
		fmt.Println("Successful")
	}
}
