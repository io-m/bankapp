package api

import (
	"net/http"
)

type Router interface {
	Get(path string, f http.HandlerFunc)
	Post(path string, f http.HandlerFunc)
	Put(path string, f http.HandlerFunc)
	Delete(path string, f http.HandlerFunc)
	Serve(addr, port string) error
}


// func GetAll(w http.ResponseWriter, r *http.Request) {
// 	customers, _ := ch.service.GetAll()
// 	w.Header().Add("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(customers); err != nil {
// 		http.Error(w, "Could not return anything", http.StatusInternalServerError)
// 	}
// }
// func (ch CustomerHandlerStruct) GetOne(w http.ResponseWriter, r *http.Request) {

// }
// func (ch CustomerHandlerStruct) Save(w http.ResponseWriter, r *http.Request) {

// }