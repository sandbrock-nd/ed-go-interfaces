package animals

import "fmt"

type Cat struct{}

func (d Cat) Meow() {
	fmt.Println("Meow")
}
