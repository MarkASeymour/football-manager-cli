package services

import (
	"encoding/json"
	"fmt"

	"github.com/markaseymour/football-manager-cli/model"
	"github.com/markaseymour/football-manager-cli/utils"
)

func GetPlayerStats(id string) ([]string, model.PlayerStatJSON, map[string]string) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?id=%s&season=2021", id)
	body := utils.MakeCall(url)

	var playStat model.PlayerStatJSON

	err3 := json.Unmarshal(body, &playStat)
	if err3 != nil {
		fmt.Println("Error unmarshalling playerstatJSON", err3)
	}

	var statList []string
	statMap := make(map[string]string)

	for _, v := range playStat.Response {
		statList = append(statList, v.Player.Name)
		statList = append(statList, v.Player.Nationality)
		statList = append(statList, v.Player.Height)
		statList = append(statList, v.Player.Weight)
		statList = append(statList, v.Statistics[0].Games.Position)
	}
	return statList, playStat, statMap

}
