package handler

import "net/http"

type CustomerHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request) 
	GetOne(w http.ResponseWriter, r *http.Request) 
	Post(w http.ResponseWriter, r *http.Request) 
	Put(w http.ResponseWriter, r *http.Request) 
	Delete(w http.ResponseWriter, r *http.Request) 
}