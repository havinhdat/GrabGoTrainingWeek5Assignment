package post

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

//go:generate mockery -name=PostInterface -inpkg
type PostInterface interface {
	GetPosts() ([]Post, error)
}
