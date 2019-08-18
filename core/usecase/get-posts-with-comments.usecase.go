package usecase

import (
	"github.com/nhaancs/GrabGoTrainingWeek5Assignment/core/entity"
)

//GetPostsWithCommentsUsecase usecase
type GetPostsWithCommentsUsecase struct {
	GetPosts GetPostsUsecase
	GetComments GetCommentsUsecase
}

// Execute GetPostsWithCommentsUsecase
func (usecase *GetPostsWithCommentsUsecase) Execute(params... interface{}) (interface{}, error) {
	postsData, err := usecase.GetPosts.Execute()
	if err != nil {
		return nil, err
	}

	commentsData, err := usecase.GetComments.Execute()
	if err != nil {
		return nil, err
	}

	// TODO: remove postsData.([]entity.Post) 
	// Go don't have generic
	posts, ok := postsData.([]entity.Post)
	if !ok {
		return nil, errTypeCasting
	}

	// TODO: remove commentsData.([]entity.Comment)
	// Go don't have generic
	comments, ok := commentsData.([]entity.Comment)
	if !ok {
		return nil, errTypeCasting
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
