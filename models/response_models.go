package models

type GetPostsResponse struct {
	Id       int64             `json:"id"`
	UserId   int64             `json:"userId"`
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	Comments []CommentResponse `json:"comments"`
}

type CommentResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}
