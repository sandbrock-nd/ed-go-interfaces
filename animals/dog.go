package animals

import "fmt"

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Bark")
}
