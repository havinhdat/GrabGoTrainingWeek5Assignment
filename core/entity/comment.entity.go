package entity

// Comment entity
type Comment struct {
	ID     int64  `xml:"id" json:"id"`
	Body   string `xml:"body" json:"body"`
	PostID int64  `xml:"postId" json:"postId"`
}