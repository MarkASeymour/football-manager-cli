package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/markaseymour/football-manager-cli/model"
	// "github.com/markaseymour/football-manager-cli/utils"
)

func GetCountriesMap() ([]string, map[string]string) {
	// config := utils.LoadConfig()

	url := "https://api-football-v1.p.rapidapi.com/v3/countries"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "566d37575dmsh34bc0d5cc7f04d8p1d34cbjsnd1db9f0f0fd1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var countriesJSON model.CountriesJSON

	err := json.Unmarshal(body, &countriesJSON)
	if err != nil {
		fmt.Println("Error unmarshalling countriesJSON %s", err)
	}

	fmt.Println(countriesJSON)
	fmt.Println("this is a result =%s", res)
	fmt.Println("this is a body =%s", body)
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