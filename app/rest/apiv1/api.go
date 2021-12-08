package apiv1

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiV1 struct {
	Router *mux.Router
	Ctx    context.Context
}

func (a *ApiV1) apiv1Handler(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	secure := a.Router.PathPrefix("/api/v1").Subrouter()
	secure.HandleFunc(path, f).Methods(method)
}

func (a *ApiV1) SetRouters() {
	a.apiv1Handler("GET", "/", a.BaseAPI)
}
