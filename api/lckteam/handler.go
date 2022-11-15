package lckteam

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gyu-young-park/lck_data_api/api/responser"
	"github.com/gyu-young-park/lck_data_api/firebaseapi"
	"github.com/gyu-young-park/lck_data_api/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) getAllTeam(res http.ResponseWriter, req *http.Request) {
	var ret LCKAllTeamResponse
	data := firebaseapi.FirebaseAppClinet.ReadTeamList()
	if data == nil {
		fmt.Println("Error: Can't get season list")
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	seasonList, ok := data["TeamList"].([]interface{})
	if !ok {
		fmt.Println("Error: Can't convert teamlist to []string")
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	for _, rawSeason := range seasonList {
		season := rawSeason.(string)
		ret.Teams = append(ret.Teams, season)
	}
	byteData, err := json.Marshal(ret)
	if err != nil {
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.ResponseJSON(res, http.StatusOK, string(byteData))
}
