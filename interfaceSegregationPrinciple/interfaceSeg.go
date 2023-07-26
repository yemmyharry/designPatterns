package main

import "fmt"

// Interface Segregation Principle
// This principle states that clients should not be forced to implement interfaces they don't use.
// Instead of one fat interface many small interfaces are preferred based on groups of methods, each one serving one submodule.

type Vehicle interface {
	Drive()
	Steer()
}

type Car struct {
}

func (c *Car) Drive() {
	fmt.Println("The car is driving")
}

func (c *Car) Steer() {
	fmt.Println("The car is steering")
}

type CDPlayer interface {
	PlayCD()
}

type CDPlayerImpl struct {
}

func (c *CDPlayerImpl) PlayCD() {
	fmt.Println("The CD player is playing a CD")
}

func main() {
	car := &Car{}
	car.Drive()
	car.Steer()

	cdPlayer := &CDPlayerImpl{}
	cdPlayer.PlayCD()
}
