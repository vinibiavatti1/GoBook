// Tests
// This module includes only testing code and uncommented examples.
// It is used only to check use cases, validate lang behaviors and others.

package main

import _ "unsafe"

func main() {

}

//go:linkname Abc
func Abc() int

//go:linkname Test
func Test() int
