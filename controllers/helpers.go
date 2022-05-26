package controllers

import "github.com/jolienai/models"

func getCommentsByPostId(id int64, comments []models.Comment) []models.CommentResponse {
	result := make([]models.CommentResponse, 0)
	for _, comment := range comments {
		if comment.PostId == id {
			result = append(result, models.CommentResponse{Name: comment.Name, Email: comment.Name, Body: comment.Body})
		}
	}
	return result
}

func createResponse(posts []models.Post, comments []models.Comment) []models.GetPostsResponse {
	response := make([]models.GetPostsResponse, 0)
	for _, post := range posts {
		comments := getCommentsByPostId(post.Id, comments)
		responseItem := models.GetPostsResponse{
			Id:       post.Id,
			UserId:   post.UserId,
			Title:    post.Title,
			Body:     post.Body,
			Comments: comments,
		}

		response = append(response, responseItem)
	}
	return response
}
