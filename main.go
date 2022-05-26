package main

import (
	"net/http"

	"github.com/jolienai/controllers"
	"github.com/jolienai/services"
)

func main() {
	service := services.NewJsonPlaceHolderService()
	controller := controllers.New(service)
	http.HandleFunc("/posts", controller.Get)
	http.ListenAndServe(":8090", nil)
}
