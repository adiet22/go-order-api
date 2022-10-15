package user

import (
	"net/http/httptest"
	"testing"

	"github.com/adiet95/go-order-api/src/database"
	"github.com/adiet95/go-order-api/src/libs"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = RepoMock{mock.Mock{}}
var service = NewService(&repo)
var ctrl = NewCtrl(service)

func TestCtrlRegister(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	pass, _ := libs.HashPassword("user12345678")
	var dataMock = database.User{Email: "user2@gmail.com", Password: "user12345678"}
	var dataMocks = database.User{Email: "user2@gmail.com", Password: pass}

	repo.mock.On("RegisterEmail", &dataMock).Return(&dataMocks, nil)

	req := httptest.NewRequest("POST", "/test/register", w.Body)

	mux.HandleFunc("/test/user", ctrl.Register)

	mux.ServeHTTP(w, req)

	var user *database.User
	respon := libs.Response{
		Data: &user,
	}

	assert.False(t, respon.IsError, "error must be false")
}
