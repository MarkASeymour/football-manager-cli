package model

type TeamsByLeagueJSON struct {
	Get        string
	Parameters struct {
		League string
		Season string
	}
	Errors  []string
	Results int
	Paging  struct {
		Current int
		Total   int
	}
	Response []struct {
		Team struct {
			Id       int
			Name     string
			Code     string
			Country  string
			Founded  int
			National bool
			Logo     string
		}
		Venue struct {
			Id       int
			Name     string
			Address  string
			City     string
			Capacity int
			Surface  string
			Image    string
		}
	}
}
