package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
)

func main() {
	fmt.Print(c.Clear)
	fmt.Println(c.Red + "Welcome to Waffles" + c.X)
	waffles := i.Ask(c.Blue + "Do you like Waffles? " + c.Red)
	if waffles == "yes" {
		if pancakes := i.Ask(c.Blue + "Do you like Pancakes? " + c.Red); pancakes == "yes" {
			if toast := i.Ask(c.Blue + "Do you like French Toast? " + c.Red); toast == "yes" {
				fmt.Println(c.Blue + "Can't wait to get a mouthful")
			} else {
				fmt.Println(c.Violet + "Well I like French Toast!")
			}
		} else {
			fmt.Println(c.Violet + "Well I like Pancakes!")
		}
	} else {
		fmt.Println(c.Violet + "Well I like Waffles!")
	}
}
