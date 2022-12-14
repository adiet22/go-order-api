package auth

import (
	"encoding/json"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
	"github.com/gin-gonic/gin"
)

type user_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(reps interfaces.AuthService) *user_ctrl {
	return &user_ctrl{reps}
}

func (u *user_ctrl) SignIn(c *gin.Context) {
	var data models.User

	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 401, true)
		c.Abort()
	}

	u.repo.Login(data).Send(c)
}

func (u *user_ctrl) Register(c *gin.Context) {
	var data *models.User

	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 401, true)
		c.Abort()
	}
	u.repo.Register(data).Send(c)
}
