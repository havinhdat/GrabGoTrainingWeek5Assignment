package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

//GetPostsWithCommentsV1 usecase
type getPostsWithCommentsV1 struct {
	GetPosts GetPostsUsecase
	GetComments GetCommentsUsecase
}

// NewGetPostsWithCommentsUsecase func
func NewGetPostsWithCommentsUsecase(getPosts GetPostsUsecase, getComments GetCommentsUsecase) GetPostsWithCommentsUsecase {
	return &getPostsWithCommentsV1 {
		GetPosts: getPosts,
		GetComments: getComments,
	}
}

// GetPostsWithComments GetPostsWithCommentsUsecase
func (usecase *getPostsWithCommentsV1) GetPostsWithComments() (*entity.PostsWithComments, error) {
	posts, err := usecase.GetPosts.GetPosts()
	if err != nil {
		return &entity.PostsWithComments{}, err
	}

	comments, err := usecase.GetComments.GetComments()
	if err != nil {
		return &entity.PostsWithComments{}, err
	}

	return combinePostsWithComments(posts, comments), nil
}

func combinePostsWithComments(posts []entity.Post, comments []entity.Comment) *entity.PostsWithComments {
	commentsByPostID := map[int64][]entity.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]entity.PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, entity.PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return &entity.PostsWithComments{Posts: result}
}
