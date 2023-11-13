package animals

import "fmt"

type Bear struct{}

func (d Bear) Growl() {
	fmt.Println("Growl")
}
