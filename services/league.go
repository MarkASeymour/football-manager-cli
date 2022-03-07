package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func GetLeagueForCountry(countryCode string) ([]string, model.CountryJSON, map[string]string) {

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
	var leagueMap = make(map[string]string)
	for _, v := range countryJSON.Response {
		leaguesNameList = append(leaguesNameList, v.League.Name)
		idString := fmt.Sprintf("%d", v.League.ID)
		leagueMap[v.League.Name] = idString
	}
	return leaguesNameList, countryJSON, leagueMap

}

func GetTeamsByLeagueId(leagueId string) ([]string, model.TeamsByLeagueJSON, map[string]string) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?league=%s&season=2021", leagueId)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", utils.LoadConfig().FootballApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var teamsByLeagueJSON model.TeamsByLeagueJSON

	err := json.Unmarshal(body, &teamsByLeagueJSON)
	if err != nil {
		fmt.Println("error unmarshalling JSON body: ", err)
	}

	var teamNamesList []string
	var teamsByLeagueMap = make(map[string]string)
	for _, v := range teamsByLeagueJSON.Response {
		teamNamesList = append(teamNamesList, v.Team.Name)
		teamsByLeagueMap[v.Team.Name] = strconv.Itoa(v.Team.Id)
	}
	return teamNamesList, teamsByLeagueJSON, teamsByLeagueMap
}
