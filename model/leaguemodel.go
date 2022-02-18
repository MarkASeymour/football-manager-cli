package model

type LeagueJSON struct {
	Get        string `json:"get"`
	Parameters struct{}
	Errors     []string
	Results    int `json:"results"`
	Paging     struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Logo string `json:"logo"`
		}
		Country struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		}
		Seasons []struct {
			Year     int    `json:"year"`
			Start    string `json:"start"`
			End      string `json:"end"`
			Current  bool   `json:"current"`
			Coverage struct {
				Fixtures struct {
					Events             bool `json:"events"`
					Lineups            bool `json:"lineups"`
					StatisticsFixtures bool `json:"statistics_fixtures"`
					StatisticsPlayers  bool `json:"statistics_players"`
				}
				Standings   bool `json:"standings"`
				Players     bool `json:"players"`
				Top_Scorers bool `json:"top_scorers"`
				Top_Assists bool `json:"top_assists"`
				Top_Cards   bool `json:"top_cards"`
				Injuries    bool `json:"injuries"`
				Predictions bool `json:"predictions"`
				Odds        bool `json:"odds"`
			}
		}
	}
}
