package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
)

func main() {
	var h = "Hello"
	fmt.Println(c.Red + h + " World!" + c.X)
	println(c.Blue + h + " World!" + c.X)
}
