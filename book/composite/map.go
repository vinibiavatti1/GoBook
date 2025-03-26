// Maps
// Maps are a collection of key-value pairs.
// In Go, maps are declared using the "map" keyword followed by the key and value types.
// The key type must be a comparable type, while the value type can be any type.
// Maps are unordered collections, meaning the order of elements is not guaranteed.
// Syntax:
//   map[<keyType>]<valueType>{<key>:<value>}

package composite

import "fmt"

// Declaring Maps
// Maps are declared using the "map" keyword followed by the key and value types.
// The key type must be a comparable type, while the value type can be any type.
func DeclaringMaps() {

	// Declaring a Map
	// When no elements are specified, the map will be empty.
	x := map[string]int{}
	fmt.Println("x:", x) // Output: x: map[]

	// Declaring a Map with Predefined Values
	// We can initialize the map with predefined key-value pairs.
	// The order of elements is not guaranteed.
	x = map[string]int{
		"A": 1,
		"B": 2,
	}
	fmt.Println("x:", x) // Output: x: map[A:1 B:2]
}

// Manipulating Maps
// We can add, access, update, and delete elements in a map.
func ManipulatingMaps() {

	// Declaring a Map
	// We will declare a map with 3 elements to be used in the examples.
	x := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	// Acessing Elements by Key
	// We access elements in a map using the key.
	// If the key is not present, the zero value of the value type is returned.
	y := x["A"]
	z := x["?"]
	fmt.Println("y:", y, "z:", z) // Output: y: 1 z: 0

	// Validating Key Existence
	// When accessing elements, we can check if the key exists in the map by looking at the second return value.
	// If the key is not present, the second return value will be false.
	v, ok := x["?"]
	fmt.Println("v:", v, "ok:", ok) // Output: v: 0 ok: false

	// Retrieving Length
	// We can get the number of key-value pairs in a map using the "len()" function.
	l := len(x)
	fmt.Println("len:", l) // Output: len: 3

	// Adding Elements
	// We can add key-value pairs to a map by assigning a value to a new key.
	x["D"] = 9
	fmt.Println("x:", x) // Output: x: map[A:1 B:2 C:3 D:9]

	// Updating Elements
	// We can update the value of an existing key by assigning a new value to it.
	x["D"] = 4
	fmt.Println("x:", x) // Output: x: map[A:1 B:2 C:3 D:4]

	// Deleting Elements
	// We can use the "delete()" builtin function to remove a key-value pair from a map.
	delete(x, "D")
	fmt.Println("x:", x) // Output: x: map[A:1 B:2 C:3]

	// Clearing a Map
	// We can use the "clear()" builtin function to remove all key-value pairs from a map.
	clear(x)
	fmt.Println("x:", x) // Output: x: map[]

	// Iterating over Maps Entries
	// We can use the for-range loop to iterate over the key-value pairs in a map.
	for k, v := range x {
		fmt.Println(k, v) // Output: A 1, B 2, C 3
	}

	// Iterating over Maps Keys
	// To iterate over the keys of a map, we can define only the key in the for-range loop.
	for k := range x {
		fmt.Println(k) // Output: A, B, C
	}

	// Iterating over Maps Values
	// To iterate over the values of a map, we can use the blank identifier "_" for the key.
	for _, v := range x {
		fmt.Println(v) // Output: 1, 2, 3
	}
}
