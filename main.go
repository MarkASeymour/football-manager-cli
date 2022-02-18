package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/markaseymour/football-manager-cli/services"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Select a country: ")
	countriesMap := services.GetCountriesMap()
	countriesNames := services.GetNamesFromMap(countriesMap)
	printOptions(countriesNames)

	text, _ := scanner.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	choice, err := strconv.ParseInt(text, 0, 64)
	if err != nil {
		fmt.Printf("Error parsing int: %s", err)
	}
	choice = choice - 1
	nameChoice := countriesNames[choice]
	choiceCode := countriesMap[nameChoice]

	leaguesForCountryList, _, leagueMap := services.GetLeagueForCountry(choiceCode)
	printOptions(leaguesForCountryList)
	leagueChoice, _ := scanner.ReadString('\n')
	leagueChoiceString := strings.TrimSuffix(leagueChoice, "\n")
	leagueChoiceNumber, err := strconv.ParseInt(leagueChoiceString, 0, 64)
	leagueChoiceIndex := leagueChoiceNumber - 1
	leagueChoiceName := leaguesForCountryList[leagueChoiceIndex]

	var leagueCode = leagueMap[leagueChoiceName]
	// fmt.Printf("league code is: %s", leagueCode)
	// leagueIndexPre, err := strconv.ParseInt(leagueChoice, 0, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// leagueIndex := leagueIndexPre - 1
	// leagueIndexString := fmt.Sprintf("%i", leagueIndex)

	teamsList, _, _ := services.GetTeamsByLeagueId(leagueCode)
	// fmt.Printf("length of league list of teams is: %d", len(teamsList))
	printOptions(teamsList)

}

func printOptions(countries []string) {
	for i, v := range countries {
		country := fmt.Sprintf("%d %s", i+1, v)
		fmt.Println(country)
	}
}
