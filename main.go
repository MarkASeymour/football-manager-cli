package main

import (
	"fmt"

	"github.com/markaseymour/football-manager-cli/services"
)

func main() {
	countryJSON := services.GetLeagueForCountry("gb")

	for _, v := range countryJSON.Response {
		fmt.Println(v.League.Name)

	}
}
