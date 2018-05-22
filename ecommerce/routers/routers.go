package routers

import (
	"net/http"

	"github.com/DiegoSantosWS/ecommerce/api"
	"github.com/gorilla/mux"
)

func Routers() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", api.MyHome).Methods("GET")
	r.HandleFunc("/api/prodocts/", api.Products).Methods("GET")
	r.HandleFunc("/api/prodocts/{slug}/{id}", api.Products).Methods("GET")

	http.ListenAndServe(":5000", r)
}
