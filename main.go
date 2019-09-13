package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

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
	Posts []PostWithComments `json:"posts"`
}
type PostWithComments struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []Comment `json:"comments,omitempty"`
}

//TODO: how to separate API logic, business logic and response format logic
func main() {
	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		// Get posts from api
		posts, err := getPosts()
		if err != nil {
			log.Println("get posts failed with error: ", err)
			writer.WriteHeader(500)
			return
		}

		// Get comments from api
		comments, err := getComments()
		if err != nil {
			log.Println("get comments failed with error: ", err)
			writer.WriteHeader(500)
			return
		}

		// Combine and return response
		postWithComments := combinePostWithComments(posts, comments)
		resp := PostWithCommentsResponse{Posts: postWithComments}
		buf, err := json.Marshal(resp)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPosts() ([]Post, error) {
	resp, err := http.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func getComments() ([]Comment, error) {
	resp, err := http.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func combinePostWithComments(posts []Post, comments []Comment) []PostWithComments {
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
