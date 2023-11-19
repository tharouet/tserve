package tserve

import "net/http"

type Router struct {
	Router *http.ServeMux
}

func NewRouter() (r *Router) {
	r.Router = http.NewServeMux()
	return
}
