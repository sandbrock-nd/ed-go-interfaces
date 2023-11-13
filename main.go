package main

import (
	"ed-go-interfaces/animals"
)

func main() {
	bear := animals.Bear{}
	cat := animals.Cat{}
	dog := animals.Dog{}
	duck := animals.Duck{}

	bear.Growl()
	cat.Meow()
	dog.Bark()
	duck.Quack()
}
