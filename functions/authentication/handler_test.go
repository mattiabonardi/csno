package function

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFunction(t *testing.T) {
	fmt.Println("Start server function on http://localhost:8080/function/authentication")
	http.HandleFunc("/function/authentication/", Handle)
	http.ListenAndServe(":8080", nil)
	fmt.Println("start")
}
