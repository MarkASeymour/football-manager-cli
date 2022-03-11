package services

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func GetCountriesMap() ([]string, map[string]string) {

	url := "https://api-football-v1.p.rapidapi.com/v3/countries"

	body := utils.MakeCall(url)

	var countriesJSON model.CountriesJSON

	err := json.Unmarshal(body, &countriesJSON)
	if err != nil {
		fmt.Println("Error unmarshalling countriesJSON %s", err)
	}

	countriesMap := make(map[string]string)
	var namesList []string
	for _, v := range countriesJSON.Response {

		countriesMap[v.Name] = v.Code
		namesList = append(namesList, v.Name)
	}
	return namesList, countriesMap

}

func GetNamesFromMap(countries map[string]string) []string {
	var namesList []string
	for k := range countries {
		namesList = append(namesList, k)
	}
	sort.Strings(namesList)
	return namesList
}
