// Error
// In Go, there isn't a try-catch mechanism, instead, errors are handled by checking the error value returned from a
// function.
// Errors in Go are values that implement the error interface.
// The error value is nil if there is no error.

package errors

import (
	"fmt"
)

// Declaring Function that Raises an Error
// The function below shows how to declare a function that returns an error.
// We can use the "errors.New" function to create a new error.
// If we need a formatted error, we can use the "fmt.Errorf" function.
// The error can be nil if there is no error.
// It returns an error value that can be used to format the error message.
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero %d/%d", x, y)
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
		fmt.Println("Error:", err) // Output: Error: division by zero 4/0
	} else {
		fmt.Println("Result:", r)
	}
}

// Wrapping an Error
// Wrapping an error means to append the current error to the new error.
// This is useful to add more context to the error.
// We can wrap an error using the "fmt.Errorf" function with the "%w" verb.
// The function below performs the other function, and wraps the error if it occurs.
func DoDivide(x, y int) (int, error) {
	r, err := Divide(x, y) // Call the function that raises an error
	if err != nil {
		return 0, fmt.Errorf("failed to divide: %w", err) // Wraps the error
	}
	return r, nil
}

// Handling Wrapped Error
// Now, when the errors is printed, it will show the wrapped error and the original error.
// This way is more recommended since it is easier to locate the error.
func HandlingWrappedError() {

	// Handling Wrapped Error
	// We will raise the error to look at the message.
	_, err := DoDivide(4, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: failed to divide: division by zero 4/0
	} else {
		fmt.Println("Successful")
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
