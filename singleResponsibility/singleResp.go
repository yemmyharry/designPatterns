package main

import "fmt"

// Single Responsibility Principle
// The Single Responsibility Principle states that a class should have only one reason to change.

type MyName struct {
	name string
}

func (s MyName) getName() string {
	return s.name
}

func (s *MyName) setName(name string) {
	s.name = name
}

func main() {
	name := MyName{}
	name.setName("John")
	fmt.Println(name.getName())

}
