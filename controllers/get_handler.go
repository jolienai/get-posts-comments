package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jolienai/models"
)

func (h *Controller) Get(w http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	postsChannel := make(chan []models.Post)
	go h.service.GetPosts(postsChannel)

	commentsChannel := make(chan []models.Comment)
	go h.service.GetComments(commentsChannel)

	response := createResponse(<-postsChannel, <-commentsChannel)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
