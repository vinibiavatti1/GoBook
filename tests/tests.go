package main

import "fmt"

func main() {
	x := NewOption(&User{ID: 1})
	if x.HasValue() {
		fmt.Println(x.Value())
	} else {
		fmt.Println("No value")
	}
}

type User struct {
	ID int
}

type Option[T any] struct {
	value *T
}

func NewOption[T any](value *T) *Option[T] {
	return &Option[T]{value: value}
}

func (o *Option[T]) HasValue() bool {
	return o.value != nil
}

func (o *Option[T]) Value() *T {
	if o.value == nil {
		panic("Option is empty")
	}
	return o.value
}

func (o *Option[T]) SetValue(value *T) {
	o.value = value
}
