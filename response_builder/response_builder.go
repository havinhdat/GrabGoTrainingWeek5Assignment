package responsebuilder

import (
	enres "../encode-response"
	getapi "../get-api"
)

type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts" xml:"Posts>Post"`
}

type PostWithComments struct {
	ID       int64            `json:"id"`
	Title    string           `json:"string"`
	Comments []getapi.Comment `json:"comments,omitempty" xml:"Comments>Comment"`
}

type ResponseBuilderImpl struct {
	render enres.EncodeResponse
}

func (builder *ResponseBuilderImpl) Build(iposts []getapi.Post, icomments []getapi.Comment) error {
	postWithComments := combinePostWithComments(iposts, icomments)
	resp := PostWithCommentsResponse{Posts: postWithComments}
	return builder.render.Encode(resp)
}

func New(irender enres.EncodeResponse) *ResponseBuilderImpl {
	return &ResponseBuilderImpl{
		render: irender,
	}
}

func combinePostWithComments(posts []getapi.Post, comments []getapi.Comment) []PostWithComments {
	commentsByPostID := map[int64][]getapi.Comment{}
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
