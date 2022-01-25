package model

// type Player struct {
// 	Id       string `json:"id"`
// 	Name     string `json:"name"`
// 	Age      string `json:"age"`
// 	Number   string `json:"number"`
// 	Position string `json:"position"`
// 	Photo    string `json:"photo"`
// }

// type Players []struct {
// 	Player []Player
// }

// type Team struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Logo string `json:"logo"`
// }

// type Details struct {
// 	Team    Team      `json:"team"`
// 	Players []Players `json:"players"`
// }

// type Response []struct {
// 	Teams Details
// }

// type Paging struct {
// 	Current int `json:"current"`
// 	Total   int `json:"total"`
// }

// type Errors []struct {
// 	ErrorArr []string
// }

// type Parameters struct {
// 	Team string `json:"team"`
// }

type SquadJSON struct {
	Get        string `json:"get"`
	Parameters struct {
		Team string `json:"team"`
	} `json:"parameters"`
	Errors []struct {
		ErrorArr []string
	} `json:"errors"`
	Results int `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Teams struct {
			Team struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			Players []struct {
				Player []struct {
					Id       string `json:"id"`
					Name     string `json:"name"`
					Age      string `json:"age"`
					Number   string `json:"number"`
					Position string `json:"position"`
					Photo    string `json:"photo"`
				}
			} `json:"players"`
		}
	} `json:"response"`
}
