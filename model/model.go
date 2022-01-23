package model

type SquadJSON struct {
	Get        string     `json:get`
	Parameters Parameters `json:parameters`
	Errors     Errors     `json:errors`
	Paging     Paging     `json:paging`
	Response   Response   `json:response`
}

type Parameters struct {
	Team string `json:team`
}

type Errors struct {
	ErrorArr []string `json:errors`
}

type Paging struct {
	Current string `json:current`
	Total   string `json:total`
}

type Response []struct {
	Teams Details
}

type Details []struct {
	Team    Team    `json:team`
	Players Players `json:players`
}
type Team struct {
	ID   string `json:id`
	Name string `json:name`
	Logo string `json:logo`
}
type Players []struct {
	Player Player
}
type Player struct {
	Id       string `json:id`
	Name     string `json:name`
	Age      string `json:age`
	Number   string `json:number`
	Position string `json:position`
	Photo    string `json:photo`
}
