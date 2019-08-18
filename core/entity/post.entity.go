package entity

// Post entity
type Post struct {
	ID    int64  `xml:"id" json:"id"`
	Title string `xml:"title" json:"title"`
}