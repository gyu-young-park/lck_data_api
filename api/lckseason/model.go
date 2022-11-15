package lckseason

type LCKAllSeasonResponse struct {
	Seasons []string `json:"seasons"`
}

func NewLCKAllSeasonResponse(seasonList []string) *LCKAllSeasonResponse {
	return &LCKAllSeasonResponse{
		Seasons: seasonList,
	}
}

type LCKTeamWithSeasonResponse struct {
	Season string   `json:"season"`
	Teams  []string `json:"teams"`
}

func NewLCKTeamWithSeasonResponse(season string, teamList []string) *LCKTeamWithSeasonResponse {
	return &LCKTeamWithSeasonResponse{
		Season: season,
		Teams:  teamList,
	}
}
