package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)
	name := input.Ask(c.Blue + "Hi, what's your name? " + c.X + c.Red)
	fmt.Println(c.X + c.Blue + "Nice to meet you " + c.X + c.Red + name + c.X)
}
