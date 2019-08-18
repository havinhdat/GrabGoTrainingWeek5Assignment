package entity

// PostWithComments entity
type PostWithComments struct {
	ID       int64     `xml:"id" json:"id"`
	Title    string    `xml:"string" json:"string"`
	Comments []Comment `xml:"comments" json:"comments,omitempty"`
}