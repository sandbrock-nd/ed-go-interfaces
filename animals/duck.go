package animals

import "fmt"

type Duck struct{}

func (d Duck) Speak() {
	fmt.Println("Quack")
}
