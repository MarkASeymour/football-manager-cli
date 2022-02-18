package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/markaseymour/football-manager-cli/model"
)

func GetCountriesMap() map[string]string {

	url := "https://api-football-v1.p.rapidapi.com/v3/countries"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "934abd1d41msh4b4711d7d89a5d8p147930jsnea416e3c7d4a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)zz

	var countriesJSON model.CountriesJSON

	err := json.Unmarshal(body, &countriesJSON)
	if err != nil {
		fmt.Println("Error unmarshalling countriesJSON %s", err)
	}
	countriesMap := make(map[string]string)
	for _, v := range countriesJSON.Response {
		countriesMap[v.Name] = v.Code
	}
	return countriesMap

}

func GetNamesFromMap(countries map[string]string) []string {
	var namesList []string
	for k := range countries {
		namesList = append(namesList, k)
	}
	return namesList
}
