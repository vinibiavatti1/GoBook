// Goto
// The "goto" statement allows jumping to a labeled statement within the same function.
// It is generally discouraged in favor of structured control flow (loops, conditionals, etc.).
// However, it can be useful for breaking out of deeply nested structures.
// Syntax:
//   goto <label>
//   <label>:

package statements

import "fmt"

// Simple Goto Example
// Demonstrates a basic jump using goto.
func SimpleGoto() {
	fmt.Println("Start") // Output: Start
	goto skip            // Jump to "skip" label
	// Unreached code
skip:
	fmt.Println("End") // Output: End
}

// Loop Simulation with Goto
// Shows how goto can be used for a manual loop-like structure.
func LoopWithGoto() {
	x := 0
repeat:
	fmt.Println("Iteration:", x)
	x++
	if x < 3 {
		goto repeat // Jumps back to "repeat" label, acting like a loop
	}
	// Output: Iteration: 0 \n Iteration: 1 \n Iteration: 2
}

// Using Goto for Error Handling
// Goto can be used to centralize error handling.
func ErrorHandlingWithGoto() {
	err := false
	if err {
		goto handleError
	}
	fmt.Println("Operation successful")
	return
handleError:
	fmt.Println("An error occurred. Handling error...")
	// Output (if err = true): An error occurred. Handling error...
}

// Breaking Out of Loops
// Goto can be used to break out of loops, unlike "break", which only exits the current loop.
func BreakingOutLoops() {
	for i := range 10 {
		if i == 3 {
			goto exit
		}
		fmt.Println(i) // Output: 0 1 2
	}
exit:
	fmt.Println("Exited nested loops") // Output: Exited nested loops
}
