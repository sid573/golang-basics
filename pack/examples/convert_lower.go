package pack

import (
	"strings"
)

type Article struct {
	Title    string
	Subtitle string
	Content  string
}

func ConvArticleList(a *Article) []string {
	p := make([]string, 0)
	result := make([]string, 0)
	tSplit := strings.Split(a.Content, " ")
	sSplit := strings.Split(a.Subtitle, " ")
	cSplit := strings.Split(a.Content, " ")
	p = append(p, tSplit...)
	p = append(p, sSplit...)
	p = append(p, cSplit...)
	for _, e := range p {
		result = append(result, strings.ToLower(e))
	}
	return result
}
