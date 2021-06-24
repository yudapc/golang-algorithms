package muxtesting

import (
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := NewMux()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
