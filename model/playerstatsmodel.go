package model

type PlayerStatJSON struct {
	Get        string
	Parameters struct {
		Id     string
		Season string
	}
	Errors  []string
	Results int
	Paging  struct {
		Current int
		Total   int
	}
	Response []struct {
		Player struct {
			Id        int
			Name      string
			Firstname string
			Lastname  string
			Age       int
			Birth     struct {
				Date    string
				Place   string
				Country string
			}
			Nationality string
			Height      string
			Weight      string
			Injured     bool
			Photo       string
		}
		Statistics []struct {
			Team struct {
				Id   string
				Name string
				Logo string
			}
			League struct {
				Id      string
				Name    string
				Country string
				Logo    string
				Flag    string
				Season  int
			}
			Games struct {
				Appearances int
				Lineups     int
				Minutes     int
				Number      int
				Position    string
				Rating      string
				Captain     bool
			}
			Substitutes struct {
				In    int
				Out   int
				Bench int
			}
			Shots struct {
				Total int
				On    int
			}
			Goals struct {
				Total    int
				Conceded int
				Assists  int
				Saves    int
			}
			Passes struct {
				Total    int
				Key      int
				Accuracy int
			}
			Tackles struct {
				Total         int
				Blocks        int
				Interceptions int
			}
			Duels struct {
				Total int
				Won   int
			}
			Dribbles struct {
				Attempts int
				Success  int
				Past     int
			}
			Fouls struct {
				Drawn     int
				Committed int
			}
			Cards struct {
				Yellow    int
				Yellowred int
				Red       int
			}
			Penalty struct {
				Won       int
				Committed int
				Scored    int
				Missed    int
				Saved     int
			}
		}
	}
}
