package order

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
	"github.com/gin-gonic/gin"
)

type order_ctrl struct {
	svc interfaces.OrderService
}

func NewCtrl(reps interfaces.OrderService) *order_ctrl {
	return &order_ctrl{svc: reps}
}

func (re *order_ctrl) GetAll(c *gin.Context) {
	v := c.Request.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(v)

	val := c.Request.URL.Query().Get("offset")
	offset, _ := strconv.Atoi(val)

	re.svc.GetAll(limit, offset).Send(c)
}

func (re *order_ctrl) Add(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	var data models.Order
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data, claim_user.(string)).Send(c)
}

func (re *order_ctrl) Update(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	email := claim_user.(string)
	val := c.Request.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	var datas models.Order
	err := json.NewDecoder(c.Request.Body).Decode(&datas)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Update(&datas, v, email).Send(c)
}

func (re *order_ctrl) Delete(c *gin.Context) {
	val := c.Request.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)

	re.svc.Delete(v).Send(c)
}

func (re *order_ctrl) Search(c *gin.Context) {
	val := c.Request.URL.Query().Get("name")
	v := strings.ToLower(val)
	re.svc.Search(v).Send(c)
}

func (re *order_ctrl) SearchId(c *gin.Context) {
	val := c.Request.URL.Query().Get("id")
	v, _ := strconv.Atoi(val)
	re.svc.SearchId(v).Send(c)
}
