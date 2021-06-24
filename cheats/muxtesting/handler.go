package muxtesting

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Home Page ")
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Category Page ")
}

func APIProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	products := GetProducts()

	responseSerialize := &ResponseProduct{
		Status: true,
		Data:   products,
	}

	json.NewEncoder(w).Encode(responseSerialize)
}

func APIProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	params := map[string]string{
		"id":     vars["id"],
		"status": r.FormValue("status"),
	}

	responseSerialize := &ResponseProduct{
		Status: true,
		Data:   params,
	}

	json.NewEncoder(w).Encode(responseSerialize)
}
