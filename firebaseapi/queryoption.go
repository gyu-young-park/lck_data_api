package firebaseapi

// collection
const (
	LCK_MATCH_COLLECTION            = "lck_match"
	LCK_TEAM_WITH_SEASON_COLLECTION = "lck_team_with_season"
	LCK_ALL_SEASONS                 = "lck_seasons"
	LCK_ALL_TEAMS                   = "lck_teams"
)

// document
const (
	LCK_ALL_SEASONS_DOC = "seasons"
	LCK_ALL_TEAMS_DOC   = "teams"
)

type ReadMatchQueryOption struct {
	Season      string
	Team        string
	Result      string
	SortOpt     string
	PublishedAt string
	Limit       string
}

func NewReadMatchQueryOption(season, team, result, sortOpt, publishedAt, limit string) *ReadMatchQueryOption {
	return &ReadMatchQueryOption{
		Season:      season,
		Team:        team,
		Result:      result,
		SortOpt:     sortOpt,
		PublishedAt: publishedAt,
		Limit:       limit,
	}
}

type ReadTeamWithSeasonQueryOption struct {
	Season string
	Team   string
}

func NewReadTeamWithSeasonQueryOption(season, team string) *ReadTeamWithSeasonQueryOption {
	return &ReadTeamWithSeasonQueryOption{
		Season: season,
		Team:   team,
	}
}