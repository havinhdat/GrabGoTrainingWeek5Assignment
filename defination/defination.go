package defination

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}
type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts" xml:"Posts>Post"`
}
type PostWithComments struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []Comment `json:"comments,omitempty" xml:"Comments>Comment"`
}
