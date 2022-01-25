package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/markaseymour/football-manager-cli/model"
)

func main() {

	url := "https://api-football-v1.p.rapidapi.com/v3/players/squads?team=38"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "934abd1d41msh4b4711d7d89a5d8p147930jsnea416e3c7d4a")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var squadJSON model.SquadJSON

	err := json.Unmarshal(body, &squadJSON)
	if err != nil {
		fmt.Println("error unmarshalling JSON body: ", err)
	}

	responseArray := squadJSON
	// fmt.Println(responseArray)
	// response := responseArray.Response
	// details := response[0]
	// teams := details.Teams
	// team := teams[0]
	// fmt.Println(string(body))

	fmt.Printf("%+v\n", responseArray)

	// fmt.Println(res)

	// fmt.Println("\n******************************************")

	// fmt.Println(string(body))

}
