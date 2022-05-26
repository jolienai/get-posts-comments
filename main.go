package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	PostId int64  `json:"postId"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type GetPostsResponse struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Comment string `json:"comment"`
}

func main() {

	http.HandleFunc("/posts", getPostsHandler)
	http.ListenAndServe(":8090", nil)

}

func getPostsHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	postsChannel := make(chan []Post)
	go getPosts(postsChannel)

	commentsChannel := make(chan []Comment)
	go getComments(commentsChannel)

	response := CreateGetPostResponse(<-postsChannel, <-commentsChannel)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)

}

func CreateGetPostResponse(posts []Post, comments []Comment) []GetPostsResponse {
	response := make([]GetPostsResponse, 0)
	for _, post := range posts {
		comment := getCommentByPostId(post.Id, comments)
		responseItem := GetPostsResponse{
			Id:      post.Id,
			UserId:  post.UserId,
			Title:   post.Title,
			Body:    post.Body,
			Comment: comment.Body,
		}

		response = append(response, responseItem)
	}
	return response
}

func getCommentByPostId(id int64, comments []Comment) Comment {
	for _, comment := range comments {
		if comment.PostId == id {
			return comment
		}
	}
	return Comment{}
}

func getComments(response chan<- []Comment) {
	comments := make([]Comment, 0)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &comments)
	if err != nil {
		panic(err)
	}
	response <- comments
}

func getPosts(response chan<- []Post) {
	posts := make([]Post, 0)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.Body == nil {
		panic(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(body, &posts)
	if jsonErr != nil {
		panic(err)
	}
	response <- posts
}
