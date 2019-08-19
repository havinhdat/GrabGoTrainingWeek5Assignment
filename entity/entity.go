package entity

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

func CombinePostWithComments(posts []Post, comments []Comment) []PostWithComments {
	commentsByPostID := map[int64][]Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}
