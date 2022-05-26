package controllers

import "github.com/jolienai/services"

type Controller struct {
	service services.JsonPlaceHolderService
}

func New(service services.JsonPlaceHolderService) *Controller {
	return &Controller{
		service: service,
	}
}
