package main

import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	one := `
 -------
|       |
|   •   |
|       |
 -------`
	two := `
 -------
| •     |
|       |
|     • | 
 -------`
	three := `
 -------
| •     |
|   •   |
|     • |
 -------`
	four := `
 -------
| •   • |
|       |
| •   • |
 -------`
	five := `
 -------
| •   • |
|   •   |
| •   • |
 -------`
	six := `
 -------
| •   • |
| •   • |
| •   • |
 -------`

	sides := []string{one, two, three, four, five, six}

	roll := rand.Intn(6)

	fmt.Println(c.Clear + c.Rc() + sides[roll])
}
