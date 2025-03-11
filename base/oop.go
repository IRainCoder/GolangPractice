/*
面向对象编程
*/
package main 

import (
	"fmt"
)

type Animal interface {
	EatFood()
	Play()
}


type Dog struct {
	Name string
	Category string
}

func (d Dog) EatFood() {
	fmt.Println(d.Name, "is eating food")
}

func (d Dog) Play() {
	fmt.Println(d.Name, "is playing")
}

type Cat struct {
	Name string
	Category string
}

func (c Cat) EatFood() {
	fmt.Println(c.Name, "is eating food")
}

func (c Cat) Play() {
	fmt.Println(c.Name, "is playing")
}

func animal(animals ...Animal) {
	for _, v := range animals {
		if _, ok := v.(Animal); ok {
			v.EatFood()
			v.Play()
		}
	}
}


func main() {
	dog := Dog{
		Name: "dog1",
		Category: "dog",
	}
	cat := Cat{
		Name: "cat1",
		Category: "cat",
	}
	animal(dog, cat)
}





