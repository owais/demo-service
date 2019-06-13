package red

import "fmt"

type Listener struct{}

func (l Listener) Start() {
	fmt.Println("starting Red")
}
