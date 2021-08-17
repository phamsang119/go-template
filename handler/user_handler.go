package handler

import (
	"game-api/constant"
	"game-api/entity"
	"game-api/service"
	"game-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//UsersHandler struct defines the dependencies that will be used
type UsersHandler struct {
	us service.UserService
}

// NewUsers UsersHandler constructor
func NewUsers(us service.UserService) *UsersHandler {
	return &UsersHandler{
		us: us,
	}
}

func (s *UsersHandler) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Reject(c, http.StatusBadRequest, constant.MissingInformation, constant.GetMessageFromCode(constant.MissingInformation), nil)
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		utils.Reject(c, http.StatusBadRequest, constant.MissingInformation, constant.GetMessageFromCode(constant.MissingInformation), nil)
		return
	}
	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		utils.Reject(c, http.StatusInternalServerError, constant.ServerError, constant.GetMessageFromCode(constant.ServerError), nil)
		return
	}
	utils.Resolve(c, http.StatusCreated, newUser.PublicUser())
}

func (s *UsersHandler) GetUsers(c *gin.Context) {
	users := entity.Users{} //customize user
	var err error
	users, err = s.us.GetUsers()
	if err != nil {
		utils.Reject(c, http.StatusInternalServerError, constant.ServerError, constant.GetMessageFromCode(constant.ServerError), nil)
		return
	}
	utils.Resolve(c, http.StatusCreated, users.PublicUsers())
}

func (s *UsersHandler) GetUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		utils.Reject(c, http.StatusBadRequest, constant.MissingInformation, constant.GetMessageFromCode(constant.MissingInformation), nil)
		return
	}
	user, err := s.us.GetUser(userId)
	if err != nil {
		utils.Reject(c, http.StatusInternalServerError, constant.ServerError, constant.GetMessageFromCode(constant.ServerError), nil)
		return
	}
	utils.Resolve(c, http.StatusOK, user.PublicUser())
}
