package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"os"
	"strconv"
)

func main() {
	fmt.Print(c.Clear)

	var max int

	if len(os.Args) < 2 {
		max, err := strconv.Atoi(i.Ask(c.Blue + "What is the max color?" + c.Red))
		if err != nil {
			panic(err)
		}
	} else {
		max, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf(c.Blue+"Multiplication tables for %s", max)

	for i := 0; i < 12; i++ {
		fmt.Printf(c.Blue+"%s x %s = %s", i, max, i*max)
	}
}
