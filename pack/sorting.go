package pack

type Player struct {
	Name  string
	Score int
}

type Players struct {
	P        []Player
	SortType string
}

func (p *Players) Len() int {
	return len(p.P)
}

func (p *Players) Less(i, j int) bool {
	if p.SortType == "Name" {
		return p.P[i].Name < p.P[j].Name
	}
	if p.SortType == "Score" {
		return p.P[i].Score < p.P[j].Score
	}
	return false
}

func (p *Players) Swap(i, j int) {
	p.P[i], p.P[j] = p.P[j], p.P[i]
}
