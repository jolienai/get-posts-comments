package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jolienai/models"
)

type JsonPlaceHolderService interface {
	GetPosts(response chan<- []models.Post)
	GetComments(response chan<- []models.Comment)
}

type service struct{}

func NewJsonPlaceHolderService() *service {
	return &service{}
}

const baseUrl = "https://jsonplaceholder.typicode.com"

func (s *service) GetPosts(response chan<- []models.Post) {
	posts := make([]models.Post, 0)
	resp, err := http.Get(baseUrl + "/posts")
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

func (s *service) GetComments(response chan<- []models.Comment) {
	comments := make([]models.Comment, 0)
	resp, err := http.Get(baseUrl + "/comments")
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
