package main

import (
	"ed-go-interfaces/animals"
)

func speak(animal animals.Animal) {
	animal.Speak()
}

func main() {
	bear := animals.Bear{}
	cat := animals.Cat{}
	dog := animals.Dog{}
	duck := animals.Duck{}

	speak(bear);
	speak(cat)
	speak(dog)
	speak(duck)
}
