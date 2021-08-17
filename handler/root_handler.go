package handler

import "food-app/service"

type RootHandler struct {
	service *service.Service
	AuthenHandler
	UsersHandler
}

func NewHandler(service *service.Service) *RootHandler {
	return &RootHandler{
		service: service,
		AuthenHandler: AuthenHandler{
			us: *service.UserService,
		},
		UsersHandler: UsersHandler{
			us: *service.UserService,
		},
	}
}
