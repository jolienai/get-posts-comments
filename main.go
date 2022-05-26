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

type PostResponse struct {
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

	response := CreateResponse(<-postsChannel, <-commentsChannel)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)

}

func CreateResponse(posts []Post, comments []Comment) []PostResponse {

	response := make([]PostResponse, 0)

	for _, p := range posts {
		c := getCommentByPostId(p.Id, comments)
		r := PostResponse{
			Id:      p.Id,
			UserId:  p.UserId,
			Title:   p.Title,
			Body:    p.Body,
			Comment: c.Body,
		}

		response = append(response, r)
	}

	return response
}

func getCommentByPostId(id int64, comments []Comment) Comment {
	for _, c := range comments {
		if c.PostId == id {
			return c
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
