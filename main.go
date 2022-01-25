package main

import (
	"fmt"

	"github.com/markaseymour/football-manager-cli/services"
)

func main() {
	var teamID string = "48"
	squadJSON := services.RunSquad(teamID)

	response := squadJSON.Response

	teamAndSquad := response[0]

	teamName := teamAndSquad.Team.Name
	squad := teamAndSquad.Players

	fmt.Println(teamName)
	fmt.Println("\n ")
	fmt.Println("CURRENT ROSTER:")

	for _, v := range squad {
		fmt.Println("\n**************")

		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Age: %v\n", v.Age)
		fmt.Printf("Number: %v\n", v.Number)
		fmt.Printf("Position: %s\n", v.Position)

		fmt.Println("\n**************")

	}

}
