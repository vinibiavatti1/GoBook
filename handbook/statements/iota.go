// Iota
// The "iota" identifier is used in constant declarations to simplify
// the definition of incrementing numbers.
// Since it can be used in expressions, it provides more flexibility
// than simple enumerations.
// The first value of "iota" is always zero.
// Every time a "const" block is declared, "iota" is reset to 0 and
// increments by one for each new constant in that block.

package statements

// Defining Successive Values
// Iota starts at zero (0) and increments automatically.
const (
	A1 = iota // 0
	B1        // 1
	C1        // 2
)

// Skipping Values
// The "_" identifier can be used to skip iota values.
const (
	_  = iota // 0 (skipped)
	B2        // 1
	C2        // 2
)

// Iota with Expressions
// We can apply expressions to modify the generated sequence.
const (
	A3 = iota + 1 // 1
	B3            // 2
	C3            // 3
)
const (
	A4 = iota * 2 // 0
	B4            // 2
	C4            // 4
)
const (
	A5 = 1 << iota // 1
	B5             // 2
	C5             // 4
	D5             // 8
)

// Using Different Data Types
// The first constant in the block can explicitly define a type,
// which all following constants will inherit.
const (
	A6 float64 = iota + 0.5 // 0.5
	B6                      // 1.5
	C6                      // 2.5
)

// Iota with Literals
// When a constant is manually assigned a value, iota continues incrementing
// but does not apply to the manually assigned constants.
const (
	A7 = iota // 0
	B7        // 1
	C7 = 999  // 999 (iota is ignored here)
	D7        // 999
	E7 = iota // 4 (iota continues incrementing)
	F7        // 5
)
