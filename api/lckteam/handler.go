package lckteam

import (
	"net/http"

	"github.com/gyu-young-park/lck_data_api/api/responser"
	"github.com/gyu-young-park/lck_data_api/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) getAllTeam(res http.ResponseWriter, req *http.Request) {
	data, err := h.repo.Get(string(repository.ALL_TEAM_WITH_SEASON))
	if err != nil {
		responser.Response(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.Response(res, http.StatusOK, data)
}
