// Struct
// Structs are composite data types that allow grouping different data together into a single unit.

package datatypes

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Defining a Struct
// We use the "struct" keyword to define structs.
// The struct below has two public fields: Name and Surname.
type Person struct {
	Name    string // Public
	surname string // Private (Not capitalized)
}

// Composition
// Composition is a type of relationship where one struct can include another as a field.
// Here, the "Employee" struct includes a pointer to the "Person" struct.
// This allows the "Employee" to access the "Person" fields, along with its own additional field, like the salary.
// We can omit the name of the composition to use the same name of the type.
type Employee struct {
	*Person // Pointer to Person, allowing direct access to Name and Surname fields
	Salary  float64
}

// Constructors
// Constructors are methods responsible for creating struct instances.
// Usually, constructors have their name with the "New" prefix.
func NewEmployee(name, surname string, salary float64) *Employee {
	return &Employee{
		Person: &Person{
			Name:    name,
			surname: surname,
		},
		Salary: salary,
	}
}

// Methods
// Methods associate functions with a struct.
// In this example, the method "FullName" is associated with the "Employee" struct and takes a pointer receiver (e *Employee).
// The other "Surname" nmethod is associated with the "Person" struct to give access to the "surname" field.
func (e *Employee) FullName() string {
	return e.Name + e.surname
}
func (p *Person) Surname() string {
	return p.surname
}

// Using Structs
// This function demonstrates how to create an instance of a struct, access its fields, mutate data, and call methods.
func UsingStructs() {

	// Creating an Instance
	// Creating an instance of "Employee", initializing the "Person" field using a pointer.
	employee1 := &Employee{
		Person: &Person{
			Name:    "John",
			surname: "Duo",
		},
		Salary: 2000.0,
	}

	// Creating an Instance Using Constructor Method
	// We can also use the constructor method to create an instance of Employee
	employee2 := NewEmployee("John", "Duo", 2000.0)

	// Accessing Public Data
	// Embedded fields can be accessed directly from the "Employee" struct.
	name1 := employee1.Person.Name // Accessing emdebbed
	name2 := employee1.Name        // Accessing directly
	fmt.Println(name1, name2)      // Output: John John

	// Calling Methods
	// Calling the "FullName" method associated to the employee.
	fullname := employee2.FullName()
	fmt.Println("fullname:", fullname) // Output: fullname: JohnDuo

	// Acessing Private Data
	// Since the "Surname" field is private, it can be accessed only by public methods
	surname1 := employee1.Person.Surname() // Accessing emdebbed
	surname2 := employee1.Surname()        // Accessing directly
	fmt.Println(surname1, surname2)        // Output: Duo Duo

	// Mutating Data
	// Changing the "Notes" field of the "Employee" struct.
	employee1.Salary = 2500.0
	fmt.Println("Salary:", employee1.Salary) // Output: Salary: 2500.0

	// Anonymous Structs
	// Anonymous structs are structs without a name.
	// If we don't need to reuse a struct, we can define and create an instance at the same time.
	person := struct {
		Name string
	}{
		Name: "John Duo",
	}
	fmt.Println("person:", person) // Output: person: {John Duo}

	// Local Structs
	// We can define a struct inside a function.
	// This struct will only be available inside the function.
	type user struct {
		Name string
	}

	// Creating an Instance of Local Struct
	// The local struct is only available inside the function
	_user := &user{
		Name: "John Duo",
	}
	fmt.Println("otherPerson:", _user) // Output: otherPerson: {John Duo}
}

// Struct Tags
// Struct tags are used to attach metadata to struct fields.
// They are commonly used to provide additional information to the encoding/json package.
// Syntax: `key:"value"`
type Customer struct {
	Id int `json:"key"`
}

// Using Struct Tags
// This function demonstrates how to use struct tags to provide metadata to fields.
func UsingStructTags() {

	// Creating an Instance
	customer := Customer{
		Id: 1,
	}

	// Parsing Struct Tags (JSON)
	// Tags are using by libraries as meta information for some operations.
	// For example, in json library, the tag "json" is used to define the key name in the JSON output.
	jsonData, _ := json.Marshal(customer)
	fmt.Println(string(jsonData)) // Output: {"key":1}

	// Accessing Struct Tags
	// Struct tags can be accessed using reflection.
	// The "reflect" package is used to inspect the struct fields and their tags.
	val := reflect.TypeOf(customer).Field(0).Tag.Get("json")
	fmt.Println(val) // Output: key
}
