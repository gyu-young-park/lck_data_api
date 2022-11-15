package lckseason

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

func (h *Handler) getAllSeason(res http.ResponseWriter, req *http.Request) {
	var ret LCKAllSeasonResponse
	data := firebaseapi.FirebaseAppClinet.ReadSeasonList()
	if data == nil {
		fmt.Println("Error: Can't get season list")
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	seasonList, ok := data["SeasonList"].([]interface{})
	if !ok {
		fmt.Println("Error: Can't convert seasonList to []string")
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	for _, rawSeason := range seasonList {
		season := rawSeason.(string)
		ret.Seasons = append(ret.Seasons, season)
	}
	byteData, err := json.Marshal(ret)
	if err != nil {
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.ResponseJSON(res, http.StatusOK, string(byteData))
}

func (h *Handler) getTeamWithSeason(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	option := firebaseapi.NewReadTeamWithSeasonQueryOption(
		query.Get("season"),
		query.Get("team"),
	)
	fmt.Println(option)

	var ret []*LCKTeamWithSeasonResponse
	data := firebaseapi.FirebaseAppClinet.ReadSeaonTeamWithQueryOption(option)
	if len(data) == 0 {
		ret = append(ret, NewLCKTeamWithSeasonResponse("", []string{""}))
		byteData, err := json.Marshal(ret)
		if err != nil {
			responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
			return
		}
		responser.ResponseJSON(res, http.StatusBadRequest, string(byteData))
		return
	}

	for _, mapper := range data {
		lckSeason, _ := mapper["Season"].(string)
		lckTeams := mapper["TeamList"].([]interface{})
		var teams []string
		for _, iTeam := range lckTeams {
			team := iTeam.(string)
			teams = append(teams, team)
		}
		ret = append(ret, NewLCKTeamWithSeasonResponse(lckSeason, teams))
	}
	byteData, err := json.Marshal(ret)
	if err != nil {
		responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.ResponseJSON(res, http.StatusOK, string(byteData))
	// data, err := h.repo.Get(string(repository.ALL_SEASON_WITH_TEAM))
	// if err != nil {
	// 	responser.ResponseJSON(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
	// 	return
	// }
	// responser.ResponseJSON(res, http.StatusOK, data)
}
