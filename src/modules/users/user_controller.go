package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")

	v := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)

	val := r.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	result := re.svc.FindEmail(claim_user.(string), limit, offset)
	result.Send(w)
}

func (re *user_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return
	}
	re.svc.Add(&data).Send(w)
}

func (re *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return
	}

	re.svc.Update(&data, claim_user.(string)).Send(w)
}

func (re *user_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("email")
	re.svc.Delete(val).Send(w)
}

func (re *user_ctrl) Search(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("email")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(w)
}

func (re *user_ctrl) SearchName(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.SearchName(v).Send(w)
}
