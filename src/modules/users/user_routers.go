package users

import (
	"github.com/adiet95/go-order-api/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.CheckAuth(ctrl.GetAll)).Methods("GET")
	route.HandleFunc("", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Add))).Methods("POST")
	route.HandleFunc("", middleware.CheckAuth(ctrl.Update)).Methods("PUT")
	route.HandleFunc("", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Delete))).Methods("DELETE")
	route.HandleFunc("/detail", middleware.CheckAuth(middleware.CheckAuthor(ctrl.Search))).Methods("GET")
	route.HandleFunc("/search", middleware.CheckAuth(middleware.CheckAuthor(ctrl.SearchName))).Methods("GET")

}
