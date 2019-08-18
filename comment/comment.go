package comment

type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

//go:generate mockery -name=CommentInterface -inpkg
type CommentInterface interface {
	GetComments() ([]Comment, error)
}
