package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.name, h.age)
}

type action struct {
	human     // вместо human Human
	isSolving bool
}

func main() {
	action := action{human{"Bob", 30}, true}
	fmt.Println(action.name, action.age) // вместо action.human.name и action.human.age

	fmt.Println(action.isSolving)

	action.greet() // вместо action.human.greet()
}
