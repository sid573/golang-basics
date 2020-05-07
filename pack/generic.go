package pack

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

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

// ExecLs to ls with hidden files
func ExecLs() {
	cmd := exec.Command("sudo", "ls", "-a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
