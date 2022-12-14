package healthy

import (
	"net/http"

	"github.com/gyu-young-park/lck_data_api/api/responser"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) health(res http.ResponseWriter, req *http.Request) {
	responser.ResponseJSON(res, http.StatusOK, "Healthy check success!\n")
}
