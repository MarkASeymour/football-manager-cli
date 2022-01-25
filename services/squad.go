package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func RunSquad(teamID string) model.SquadJSON {

	config := utils.LoadConfig()

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players/squads?team=%s", teamID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.FootballApiKey)

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var squadJSON model.SquadJSON

	err := json.Unmarshal(body, &squadJSON)
	if err != nil {
		fmt.Println("error unmarshalling JSON body: ", err)
	}

	return squadJSON
}
