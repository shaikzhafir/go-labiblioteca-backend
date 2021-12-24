package domain

type Book struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	ImageURL    string `json:"imageURL"`
}
