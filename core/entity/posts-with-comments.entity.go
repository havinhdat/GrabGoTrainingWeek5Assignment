package entity

// PostsWithComments entity
type PostsWithComments struct {
	Posts []PostWithComments `xml:"posts" json:"posts"`
}