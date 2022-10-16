package routers

import (
	"errors"
	"net/http"

	"github.com/adiet95/go-order-api/src/database"
	auth "github.com/adiet95/go-order-api/src/modules/auth"
	"github.com/adiet95/go-order-api/src/modules/order"
	"github.com/adiet95/go-order-api/src/modules/users"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	auth.New(mainRoute, db)
	users.New(mainRoute, db)
	order.New(mainRoute, db)

	return mainRoute, nil
}
