package main

import "fmt"

type human struct {
	name string
	age  int
}
type action struct {
	human     // embedding: вместо human Human
	isSolving bool
}

func (person *human) setAge(age int) {
	person.age = age
}
func (h action) greet() {
	fmt.Printf("yo, my name's %s and i'm %d years old.\nwhat do you think i'm doing? %t, i'm solvin' WB tasks rn ", h.name, h.age, h.isSolving)
	// embedding: вместо h.human.name и h.human.age
}

func main() {
	action := action{human{name: "Deadinside"}, true}

	action.setAge(20) // embedding: тип action вызывает метод *human

	fmt.Println(action.name, action.age) // embedding: вместо action.human.name и action.human.age

	action.greet()
}
