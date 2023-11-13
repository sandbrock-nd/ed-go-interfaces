package animals

import "fmt"

type Bear struct{}

func (d Bear) Speak() {
	fmt.Println("Growl")
}
