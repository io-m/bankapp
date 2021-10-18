package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/io-m/bankapp/internal"
	"github.com/io-m/bankapp/internal/errors"
	service "github.com/io-m/bankapp/src/service/customer"
)

type customerHandler struct {
	service service.CustomerServicer
}

func NewCustomerHandler(service service.CustomerServicer) CustomerHandler {
	return &customerHandler{
		service: service,
	}
}


func (ch *customerHandler) GetAll(w http.ResponseWriter, r *http.Request)  {
	customers, err := ch.service.GetAll()
	if err != nil {
		internal.ResponseError(w, errors.NewError(err.Error(), http.StatusInternalServerError))
		return
	}
	internal.ResponseJSON(w, customers, http.StatusOK)
}
func (ch *customerHandler) GetOne(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errors.NewError("Bad Request", http.StatusBadRequest)
	}
	customer, err := ch.service.GetOne(id)
	if err != nil {
		internal.ResponseError(w, errors.NewError(err.Error(), http.StatusInternalServerError))
		return
	}
	internal.ResponseJSON(w, customer, http.StatusOK)
}
func (ch *customerHandler) Post(w http.ResponseWriter, r *http.Request)  {}
func (ch *customerHandler) Put(w http.ResponseWriter, r *http.Request)  {}
func (ch *customerHandler) Delete(w http.ResponseWriter, r *http.Request)  {}