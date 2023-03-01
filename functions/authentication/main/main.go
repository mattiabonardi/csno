package main

import (
	"fmt"
	"handler/function"
	"net/http"
)

func main() {
	fmt.Println("Start server function on http://localhost:8080/function/authentication")
	http.HandleFunc("/function/authentication", function.Handle)
	http.ListenAndServe(":8080", nil)
	fmt.Println("start")
}
