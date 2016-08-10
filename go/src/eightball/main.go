package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	i "github.com/skilstak/go-input"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Print(c.Clear)
	rand.Seed(time.Now().UnixNano())
	answers := []string{"yes", "no", "maybe", "someday", "ask later", "go away", "7"}
	for {
		question := i.Ask(c.Random() + "What's your question?\n")

		if strings.HasPrefix(question, "hi") {
			fmt.Println(c.Random() + "hello!")
		} else if strings.HasSuffix(question, "please") {
			fmt.Println(c.Rc() + "Never!")
		} else {
			fmt.Println(answers[rand.Intn(len(answers))])
		}
	}
}
