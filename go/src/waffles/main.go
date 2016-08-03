package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)
	fmt.Println(c.Red + "Welcome to Waffles" + c.X)
	waffles := i.Ask(c.Blue + "Do you like Waffles? " + c.X)
	if waffles == "yes" {
		if pancakes := i.Ask(c.Blue + "Do you like pancakes? "); pancakes == "yes" {
			if toast := i.Ask(c.Blue + "Do you like french toast? "); toast == "yes" {
				fmt.Println(c.Blue + "Can't wait to get a mouthful")
			}
		}
	}
}
