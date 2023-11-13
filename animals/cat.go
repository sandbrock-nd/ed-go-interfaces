package animals

import "fmt"

type Cat struct{}

func (d Cat) Speak() {
	fmt.Println("Meow")
}
