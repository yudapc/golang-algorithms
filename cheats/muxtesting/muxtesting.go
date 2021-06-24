package muxtesting

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet).Name("HomePath")
	r.HandleFunc("/category", CategoryHandler).Methods(http.MethodGet).Name("CategoryPath")
	r.HandleFunc("/api/products", APIProductsHandler).Methods(http.MethodGet).Name("APIProducts")
	r.HandleFunc("/api/products/{id}", APIProductDetailHandler).
		Methods(http.MethodGet).
		Name("APIProductDetail")

	return r
}
