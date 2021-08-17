package handler

import "game-api/service"

type RootHandler struct {
	service *service.Service
	*AuthenHandler
	*UsersHandler
}

func NewHandler(svc *service.Service) *RootHandler {
	return &RootHandler{
		service:       svc,
		AuthenHandler: NewAuthenHandler(svc.UserService),
		UsersHandler:  NewUsers(svc.UserService),
	}
}
