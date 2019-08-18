package api

import (
	"net/http"

	comment_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/comment"
	encoder "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/encode"
	error_handler "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/errorhandler"
	post_getter "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post"
	post_comment_combiner "dat.havinh/week5-assignment/GrabGoTrainingWeek5Assignment/post-comment"
)

type PostWithCommentsResponse struct {
	Posts []post_comment_combiner.PostWithComments `json:"posts"`
}

type PostWithComments struct {
	ID       int64                    `json:"id"`
	Title    string                   `json:"string"`
	Comments []comment_getter.Comment `json:"comments,omitempty"`
}

type Api interface {
	GetPostsWithComments(writer http.ResponseWriter, request *http.Request)
}

type ApiImpl struct {
	postGetter          post_getter.PostGetter
	commentGetter       comment_getter.CommentGetter
	postCommentCombiner post_comment_combiner.PostWithCommentsCombiner
	encoder             encoder.EncodeResponse
}

func (apiImpl *ApiImpl) GetPostsWithComments(writer http.ResponseWriter, request *http.Request) {
	posts, err := apiImpl.postGetter.GetPosts()
	handleError(&error_handler.PostErrorHandler{}, err, writer)

	comments, err := apiImpl.commentGetter.GetComments()
	handleError(&error_handler.CommentErrorHandler{}, err, writer)

	postWithComments := apiImpl.postCommentCombiner.CombinePostWithComments(posts, comments)
	resp := PostWithCommentsResponse{Posts: postWithComments}

	err = apiImpl.encoder.Encode(writer, resp)
	handleError(&error_handler.PostWithCommentsErrorHandler{}, err, writer)
}

func handleError(errHandler error_handler.ErrorHandler, err error, writer http.ResponseWriter) {
	errHandler.Handle(err, writer)
}

func New(postGetter post_getter.PostGetter, commentGetter comment_getter.CommentGetter,
	postCommentCombiner post_comment_combiner.PostWithCommentsCombiner, encoder encoder.EncodeResponse) (Api, error) {
	return &ApiImpl{
		postGetter:          postGetter,
		commentGetter:       commentGetter,
		postCommentCombiner: postCommentCombiner,
		encoder:             encoder,
	}, nil
}
