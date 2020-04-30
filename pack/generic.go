package pack

import "fmt"

type Generic interface {
	String() string
	// Int() int
}

type Student struct {
	Name string
	Roll int
	NewS bool
}

type CompFunc func() string

func (c CompFunc) Print() {
	fmt.Println(c())
}

func (c *Student) String() string {
	return fmt.Sprintf("Name: %v \nRoll: %d \nNew Student: %t \n", c.Name, c.Roll, c.NewS)
}
