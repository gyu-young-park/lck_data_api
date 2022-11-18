package lckmatch

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gyu-young-park/lck_data_api/api/responser"
	"github.com/gyu-young-park/lck_data_api/api/validator"
	"github.com/gyu-young-park/lck_data_api/firebaseapi"
	"github.com/gyu-young-park/lck_data_api/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) getAllMatch(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	option := firebaseapi.NewReadMatchQueryOption(
		query.Get("season"),
		query.Get("team"),
		query.Get("winLose"),
		query.Get("sortOption"),
		query.Get("publishedAt"),
		query.Get("limit"),
	)
	fmt.Println(option)
	err := validator.Client.Validate(option)
	if err != nil {
		fmt.Println(err)
		responser.ResponseJSON(res, http.StatusBadRequest, "Error client request is not valid, please check your request is valid or not\n")
		return
	}
	// data, err := h.repo.Get(string(repository.ALL_MATCH))
	// if err != nil {
	// 	responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
	// 	return
	// }
	// responser.ResponseJSON(res, http.StatusOK, data)
	data := firebaseapi.FirebaseAppClinet.ReadMatchTeamWithQueryOption(option)
	byteData, err := json.Marshal(data)
	if err != nil {
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.ResponseJSON(res, http.StatusOK, string(byteData))
}
