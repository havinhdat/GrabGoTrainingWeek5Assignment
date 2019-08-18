package post

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type PostInterface interface {
	GetPosts() ([]Post, error)
}
