package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func RunLeague() model.LeagueJSON {

	config := utils.LoadConfig()

	url := "https://api-football-v1.p.rapidapi.com/v3/leagues"

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

	return leagueJSON

}

func GetLeagueForCountry(countryCode string) model.CountryJSON {

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

	return countryJSON

}
