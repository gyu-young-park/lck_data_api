package lckteam

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gyu-young-park/lck_data_api/repository"
)

const HTTP_ROUTER_PREFIX_LCK_TEAM = "/lck-team"

type Router struct {
	handler *Handler
}

func NewRouter(repo repository.Repository) *Router {
	return &Router{handler: NewHandler(repo)}
}

func (r *Router) Route(mux *mux.Router) {
	subRouter := mux.PathPrefix(HTTP_ROUTER_PREFIX_LCK_TEAM).Subrouter()
	subRouter.HandleFunc("", r.handler.getAllTeam).Methods(http.MethodGet)
}
