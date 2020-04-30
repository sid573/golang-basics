package main

import (
	"fmt"
	"sort"

	"learn/pack"
)

func main() {

	S1 := &pack.Student{
		Name: "Siddhant Sinha",
		Roll: 48,
		NewS: true,
	}

	S2 := &pack.Student{
		Name: "Kedar Nath",
		Roll: 34,
		NewS: false,
	}

	pack.CompFunc(S1.String).Print()
	PCheck(S2)

	arr := []pack.Player{
		pack.Player{
			Name:  "Siddhant",
			Score: 100,
		},
		pack.Player{
			Name:  "Sachin",
			Score: 99,
		},
		pack.Player{
			Name:  "Kohli",
			Score: 12,
		},
	}

	player := &pack.Players{
		P:        arr,
		SortType: "Name",
	}

	fmt.Println(player.P)
	sort.Sort(player)
	fmt.Println(player.P)

}

// PCheck is used to print
func PCheck(s pack.Generic) {
	fmt.Println(s)
}
