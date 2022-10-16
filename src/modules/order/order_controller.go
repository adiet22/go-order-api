package order

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
)

type order_ctrl struct {
	svc interfaces.OrderService
}

func NewCtrl(reps interfaces.OrderService) *order_ctrl {
	return &order_ctrl{svc: reps}
}

func (re *order_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)

	val := r.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	re.svc.GetAll(limit, offset).Send(w)
}

func (re *order_ctrl) Add(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")
	email := claim_user.(string)

	var data models.Order
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return
	}
	re.svc.Add(&data, email).Send(w)
}

func (re *order_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	claim_user := r.Context().Value("email")
	email := claim_user.(string)

	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	var datas models.Order
	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.New(err.Error(), 400, true)
		return
	}
	re.svc.Update(&datas, v, email).Send(w)
}

func (re *order_ctrl) Delete(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(w)
}

func (re *order_ctrl) Search(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(w)
}

func (re *order_ctrl) SearchId(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)
	re.svc.SearchId(v).Send(w)
}
