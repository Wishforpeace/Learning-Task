package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

var person *Person
var once sync.Once

func NewPerson(name string) *Person {
	once.Do(func() {
		person = new(Person)
		person.Name = name
	})

	return person
}
func main() {
	p1 := NewPerson("hallen1")
	p2 := NewPerson("hallen2")
	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)

}
