package function

import (
	"handler/function/handlers"
	"handler/function/managers"
	"net/http"
	"strings"
)

var routes = map[string]func(http.ResponseWriter, *http.Request){}

func init() {
	routes["/login"] = handlers.LoginHanlder()
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/login") {
		routes["/login"](w, r)
		return
	}

	managers.ThrowNotFoundError(w)
}
