// Datatypes
// Go has several built-in data types, categorized as follows:
// - Boolean
// - String
// - Integer (Int, Byte and Rune)
// - Float
// - Complex
// - Container (Array, Slice and Map)
// - Pointer (*)
// - Function
// - Struct
// - Interface
// - Channel

package datatypes

// Boolean
var _ bool = true
var _ bool = false

// String
var _ string = "Hello"

// Integer
var _ int = 1     // 32/64-bit integer (default)
var _ int8 = 1    // 8-bit: -128 to 127
var _ int16 = 1   // 16-bit: -32768 to 32767
var _ int32 = 1   // 32-bit: -2147483648 to 2147483647
var _ int64 = 1   // 64-bit: -9223372036854775808 to 9223372036854775807
var _ uint = 1    // 32/64-bit Unsigned integer (default)
var _ uint8 = 1   // 8-bit: 0 to 255
var _ uint16 = 1  // 16-bit: 0 to 65535
var _ uint32 = 1  // 32-bit: 0 to 4294967295
var _ uint64 = 1  // 64-bit: 0 to 18446744073709551615
var _ uintptr = 1 // Memory Address Type (32 or 64 bits)
var _ byte = 1    // Alias for uint8
var _ rune = 'A'  // Alias for int32 (Unicode code point)

// Float
var _ float32 = 3.14 // 32-bit float
var _ float64 = 3.14 // 64-bit float (default)

// Complex
var _ complex64 = 1 + 2i  // 32-bit real and imaginary parts
var _ complex128 = 1 + 2i // 64-bit real and imaginary parts (default)

// Container
var _ = [3]int{1, 2, 3} // Array (Fixed Length)
var _ = []int{1, 2, 3}  // Slice (Dynamic Length)
var _ = map[string]int{ // Map (Key -> Value)
	"A": 1,
	"B": 2,
}

// Pointer
var _ *int = nil

// Function
func Function(x, y int) int {
	return x + y
}

// Struct
type User struct {
	Name string
}

// Interface
type Handler interface {
	Handle()
}

// Channel
var _ chan int = make(chan int)
