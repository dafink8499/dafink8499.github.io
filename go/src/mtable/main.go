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
	var err error

	if len(os.Args) < 2 {
		max, err = strconv.Atoi(i.Ask(c.Blue + "What is the max number? " + c.Red))
	} else {
		max, err = strconv.Atoi(os.Args[1])
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf(c.Blue+"Multiplication tables for %d\n", max)

	for i := 1; i < 13; i++ {
		fmt.Printf(c.Blue+"%d x %d = %d\n", i, max, i*max)
	}
}
