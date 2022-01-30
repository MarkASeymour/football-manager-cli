package main

import (
	"fmt"

	"github.com/markaseymour/football-manager-cli/services"
)

func main() {
	leagueJSON := services.RunLeague()

	for _, v := range leagueJSON.Response {
		if v.Country.Name == "England" {
			if v.League.Name == "Premier League" {
				fmt.Println(v)
			}

		}

	}
}
