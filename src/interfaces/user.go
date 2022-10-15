package interfaces

import (
	"net/http"

	"github.com/adiet95/go-order-api/src/database"
	"github.com/adiet95/go-order-api/src/libs"
)

type UserRepo interface {
	FindByEmail(username string) (*database.User, error)
	RegisterEmail(data *database.User) (*database.User, error)
}

type UserService interface {
	Login(body database.User, w http.ResponseWriter) *libs.Response
	Register(body *database.User) *libs.Response
}
