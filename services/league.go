package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func GetLeagueInfo(leagueCode string) ([]string, map[string]int) {

	config := utils.LoadConfig()

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?id=%s", leagueCode)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.FootballApiKey)

	res, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		fmt.Println("Error retrieving api league data")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var leagueJSON model.LeagueJSON

	err := json.Unmarshal(body, &leagueJSON)
	if err != nil {
		fmt.Println("error unmarshalling JSON body: ", err)
	}
	var teamsList []string
	var teamsMap map[string]int

	for _, v := range leagueJSON.Response {
		teamsList = append(teamsList, v.League.Name)
		teamsMap[v.League.Name] = v.League.ID

	}

	return teamsList, teamsMap

}

func GetLeagueForCountry(countryCode string) ([]string, model.CountryJSON) {

	config := utils.LoadConfig()

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?code=%s", countryCode)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.FootballApiKey)

	res, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		fmt.Println("Error retrieving api league data")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var countryJSON model.CountryJSON

	err := json.Unmarshal(body, &countryJSON)
	if err != nil {
		fmt.Println("error unmarshalling JSON body: ", err)
	}
	var leaguesNameList []string
	for _, v := range countryJSON.Response {
		leaguesNameList = append(leaguesNameList, v.League.Name)
	}
	return leaguesNameList, countryJSON

}
