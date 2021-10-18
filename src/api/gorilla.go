package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/io-m/bankapp/internal/logger"
)

type gorillaRouter struct {
	router *mux.Router
}

func NewGorillaRouter() Router {
	return &gorillaRouter{
		router: mux.NewRouter(),
	}
}

func (gr *gorillaRouter) Get(path string, f http.HandlerFunc) {
	gr.router.HandleFunc(path, f).Methods(http.MethodGet)
}
func (gr *gorillaRouter) Post(path string, f http.HandlerFunc) {
	gr.router.HandleFunc(path, f).Methods(http.MethodPost)
}
func (gr *gorillaRouter) Put(path string, f http.HandlerFunc) {
	gr.router.HandleFunc(path, f).Methods(http.MethodPut)
}
func (gr *gorillaRouter) Delete(path string, f http.HandlerFunc) {
	gr.router.HandleFunc(path, f).Methods(http.MethodDelete)
}
func (gr *gorillaRouter) Serve(addr, port string) error {
	logger.Info.Printf("Server is running on %s:%s", addr, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), gr.router)
}
