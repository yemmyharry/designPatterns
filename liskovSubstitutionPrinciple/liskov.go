package main

import "fmt"

type Bird interface {
	Fly()
}

type Penguin struct{}

func (p *Penguin) Fly() {
	fmt.Println("Penguins can't fly.")
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
	fmt.Println("Sparrow is flying.")
}

func FlyInSky(b Bird) {
	b.Fly()
}

func main() {

	peng := Penguin{}
	FlyInSky(&peng)

}
