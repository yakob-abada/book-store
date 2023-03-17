package domain

type Book struct {
	Name     string `json:"name"`
	Featured bool   `json:"featured"`
	Author   string `json:"author"`
}
