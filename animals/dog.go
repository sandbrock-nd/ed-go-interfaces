package animals

import "fmt"

type Dog struct{}

func (d Dog) Bark() {
	fmt.Println("Bark")
}
