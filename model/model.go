package model

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}
type PostWithComments struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []Comment `json:"comments,omitempty"`
}
type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts"`
}
