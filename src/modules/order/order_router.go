package order

import (
	"github.com/adiet95/go-order-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/order").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.CheckAuth(ctrl.Add)).Methods("POST")
	route.HandleFunc("", middleware.CheckAuth(ctrl.Update)).Methods("PUT")
	route.HandleFunc("", middleware.CheckAuth(ctrl.Delete)).Methods("DELETE")
	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth(ctrl.Search)).Methods("GET")
	route.HandleFunc("/detail", middleware.CheckAuth(ctrl.SearchId)).Methods("GET")

}
