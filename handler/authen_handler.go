package handler

import (
	"game-api/service"
	"game-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenHandler struct {
	us service.UserService
}

// NewAuthenHandler  constructor
func NewAuthenHandler(uApp service.UserService) *AuthenHandler {
	return &AuthenHandler{
		us: uApp,
	}
}

func (au *AuthenHandler) Login(c *gin.Context) {

	utils.Resolve(c, http.StatusOK, nil)
}

func (au *AuthenHandler) Logout(c *gin.Context) {

	utils.Resolve(c, http.StatusOK, nil)
}
