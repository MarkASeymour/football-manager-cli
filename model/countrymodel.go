package model

type CountryJSON struct {
	Get        string `json:"get"`
	Parameters struct {
		Current string `json:"current"`
		Code    string `json:"code"`
		Season  string `json:"season"`
		Type    string `json:"type"`
	} `json:"parameters"`
	Errors  []string `json:"errors"`
	Results int      `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Logo string `json:"logo"`
		} `json:"league"`
		Country struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		} `json:"country"`
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
				} `json:"fixtures"`
				Standings   bool `json:"standings"`
				Players     bool `json:"players"`
				TopScorers  bool `json:"top_scorers"`
				TopAssists  bool `json:"top_assists"`
				TopCards    bool `json:"top_cards"`
				Injuries    bool `json:"injuries"`
				Predictions bool `json:"predictions"`
				Odds        bool `json:"odds"`
			}
		} `json:"seasons"`
	} `json:"response"`
}
