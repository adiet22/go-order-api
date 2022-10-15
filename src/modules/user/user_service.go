package user

import (
	"net/http"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
)

type user_service struct {
	repo interfaces.UserRepo
}
type token_response struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (u user_service) Login(body models.User, w http.ResponseWriter) *libs.Response {
	checkRegist := libs.Validation(body.Email, body.Password)
	if checkRegist != nil {
		return libs.New(checkRegist.Error(), 400, true)
	}

	user, err := u.repo.FindByEmail(body.Email)
	if err != nil {
		return libs.New("email not registered, register first", 401, true)
	}
	if !libs.CheckPass(user.Password, body.Password) {
		return libs.New("wrong password", 401, true)
	}
	token := libs.NewToken(body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.New(err.Error(), 401, true)
	}
	w.Header().Set("Access", theToken)

	return libs.New(token_response{Tokens: theToken}, 200, false)
}

func (u user_service) Register(body *models.User) *libs.Response {
	checkRegist := libs.Validation(body.Email, body.Password)
	if checkRegist != nil {
		return libs.New(checkRegist.Error(), 400, true)
	}

	hassPass, err := libs.HashPassword(body.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	body.Password = hassPass
	result, err := u.repo.RegisterEmail(body)
	if err != nil {
		return libs.New("email already registered, please login", 401, true)
	}
	return libs.New(result, 200, false)
}
