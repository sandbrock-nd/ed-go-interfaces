package animals

import "fmt"

type Duck struct{}

func (d Duck) Quack() {
	fmt.Println("Quack")
}
