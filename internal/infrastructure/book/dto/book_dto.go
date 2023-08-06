package dto

type Book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b Book) HasISBN() bool {
	return b.ISBN != ""
}
