package main

import (
	"learn/pack"
)

/*
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
	b, _ := json.Marshal(player.P)
	CtB := []byte(string(b))
	fp, _ := pack.OSCustomOpen("player.txt", CtB)
	fp.Close()

	go pack.Learn()
	fmt.Println("Main.go")
	time.Sleep(1 * time.Second)

}

// PCheck is used to print
func PCheck(s pack.Generic) {
	fmt.Println(s)
}
*/
/*
func main() {
	d, _ := os.UserCacheDir()
	fmt.Println(d)
	pack.CustomCopyBuffer()
}
*/

func main() {
	c := &pack.Person{
		Name:  "Siddhant",
		Limit: 23,
	}

	pack.Random(c)
}
