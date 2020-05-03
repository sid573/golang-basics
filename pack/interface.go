package pack

import (
	"fmt"
)

type Corona interface {
	test() bool
}

type Person struct {
	Name  string
	Limit int
}

func (p *Person) test() bool {
	if p.Limit > 30 {
		return true
	}
	return false
}

// Random is a function to interface
func Random(c Corona) {
	fmt.Println(c.test())
}

/*
This whole concept is so much fun try learning it
I started with io interfaces and got confused but
I want to quote something
"Confusion at its peak is real understanding"
*/
