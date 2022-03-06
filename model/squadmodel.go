package model

type SquadJSON struct {
	Get        string `json:"get"`
	Parameters struct {
		Team string `json:"team"`
	} `json:"parameters"`
	Errors  []string `json:"errors"`
	Results int      `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"team"`
		Players []struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Age      int    `json:"age"`
			Number   int    `json:"number"`
			Position string `json:"position"`
			Photo    string `json:"photo"`
		} `json:"players"`
	} `json:"response"`
}
