package types

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      Author `json:"author"`
}
