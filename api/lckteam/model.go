package lckteam

type LCKAllTeamResponse struct {
	Teams []string `json:"teams"`
}

func NewLCKAllTeamResponse(teamList []string) *LCKAllTeamResponse {
	return &LCKAllTeamResponse{
		Teams: teamList,
	}
}
