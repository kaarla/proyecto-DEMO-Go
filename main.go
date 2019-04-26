package main

import (
	//"fmt"
	"net/http"
    //"html/template"
	"github.com/gorilla/mux"
    "github.com/kaarla/proyecto-DEMO/app/controller"
    
)

//var Products Product[]

func main() {
	r := newRouter()
    //http.HandleFunc("./assets/", controller.GetProductsHandler)
	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
    r.HandleFunc("/products", controller.GetProductsHandler).Methods("GET")
	return r
}
