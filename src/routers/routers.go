package routers

import (
	"errors"
	"net/http"

	"github.com/adiet95/go-order-api/src/database"
	user "github.com/adiet95/go-order-api/src/modules/user"
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

	user.New(mainRoute, db)

	return mainRoute, nil
}
