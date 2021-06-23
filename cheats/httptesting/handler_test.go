package httptesting

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHello(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
