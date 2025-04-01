// Error
// In Go, there isn't a try-catch mechanism, instead, errors are handled by checking the error value returned from a
// function.
// Errors in Go are values that implement the error interface.
// The error value is nil if there is no error.

package errors

import "fmt"

// Declaring Function that can Raises an Error
// The function below shows how to declare a function that returns an error.
// The error can be nil if there is no error.
// We can use the "fmt.Errorf" function to create an error.
// We can also use the "errors" package to create errors.
func ValidateAge(age int) (int, error) {
	if age > 0 {
		return age, nil
	}
	return 0, fmt.Errorf("invalid age: %d", age)
}

// Handling Error
// The function below shows how to handle an error.
func HandlingError() {

	// Performing Successfully
	// If the function performs successfully, the error value will be nil.
	age, err := ValidateAge(25)
	fmt.Println("Age:", age, "Error:", err) // Output: Age: 25 Error: <nil>

	// Performing Unsuccessfully
	// If the function fails, the error value will not be nil.
	age, err = ValidateAge(-1)
	fmt.Println("Age:", age, "Error:", err) // Output: Age: 0 Error: invalid age: -1

	// Handling Errors
	// We can check if the error value is nil or not.
	// If the error value is not nil, we can handle the error.
	age, err = ValidateAge(-1)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: invalid age: -1
	} else {
		fmt.Println("Age is valid:", age)
	}
}

// Declaring Function that can Raises Multiple Errors
// We can use a slice of errors to raise multiple errors.
// The function below shows how to declare a function that returns multiple errors.
func ValidatePassword(p string) []error {
	var errors []error
	if p == "" {
		errors = append(errors, fmt.Errorf("password is empty"))
	}
	if len(p) < 8 {
		errors = append(errors, fmt.Errorf("password is too short"))
	}
	return errors
}

// Handling Multiple Errors
// The function below shows how to handle multiple errors.
func HandlingMultipleErrors() {

	// Handling Multiple Errors
	// We can check if the error slice is empty or not.
	// If the error slice is not empty, we can handle the errors.
	passwordErrors := ValidatePassword("123")
	if len(passwordErrors) > 0 {
		for _, err := range passwordErrors {
			fmt.Println("Error:", err) // Output: Error: password is too short
		}
	} else {
		fmt.Println("Password is valid")
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
func HandlingCustomError() {

	// Handling Custom Error
	// We will call the function above to raise the custom error.
	_, err := Request("/admin")
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: Code: 403, Message: Forbidden
	} else {
		fmt.Println("Request is successful")
	}
}
