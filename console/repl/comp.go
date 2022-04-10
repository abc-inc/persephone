package repl

type Item struct {
	View    string `json:"view"`
	Content string `json:"content"`
}

func (i Item) String() string {
	return i.View
}

type CompFunc func(str string) []Item
