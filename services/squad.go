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

func GetSquad(teamID string) ([]string, model.SquadJSON, map[string]string) {

	config := utils.LoadConfig()

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players/squads?team=%s", teamID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.FootballApiKey)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var squadJSON model.SquadJSON

	err3 := json.Unmarshal(body, &squadJSON)
	if err3 != nil {
		fmt.Println("error unmarshalling JSON body: ", err3)
	}

	var namesList []string
	var namesCodeMap = make(map[string]string)

	for _, v := range squadJSON.Response {
		for _, t := range v.Players {

			namesList = append(namesList, t.Name)
			stringID := strconv.Itoa(t.Id)
			namesCodeMap[t.Name] = stringID
		}
	}

	return namesList, squadJSON, namesCodeMap
}
