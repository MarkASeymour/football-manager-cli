package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/logrusorgru/aurora"

	"github.com/markaseymour/football-manager-cli/services"
)

var prevChoices map[int]string = make(map[int]string)

func main() {

	// json := services.GetPlayerStats("256")
	// fmt.Println(json)

	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

}

type model struct {
	choices      []string
	cursor       int
	page         int
	namesCodeMap map[string]string
	allChoices   []string
	allChunks    [][]string
	currentChunk int
	isPlayerStat bool
	// prevChoices map[int]string
}

func initialModel() model {
	countryNames, namesMap := services.GetCountriesMap()
	initialChunks := chunkSlice(countryNames, 8)
	return model{
		page:         0,
		namesCodeMap: namesMap,
		currentChunk: 0,
		allChoices:   countryNames,
		allChunks:    initialChunks,
		choices:      initialChunks[0],
		isPlayerStat: false,
	}
}
func newPageModel(newList []string, newMap map[string]string, newPageNum int, isPlayStat bool) model {

	allChunksNew := chunkSlice(newList, 8)
	return model{
		page:         newPageNum,
		namesCodeMap: newMap,
		currentChunk: 0,
		allChoices:   newList,
		allChunks:    allChunksNew,
		choices:      allChunksNew[0],
		isPlayerStat: isPlayStat,
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selection := m.choices[m.cursor]

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			if m.cursor == 0 {
				if m.currentChunk > 0 {
					m.currentChunk = m.currentChunk - 1
					m.choices = m.allChunks[m.currentChunk]
				}

			}

		case "down", "j":

			if m.cursor <= len(m.choices)-1 {
				m.cursor++
			}

			if m.cursor == len(m.choices) && m.currentChunk != len(m.allChunks)-1 {

				m.currentChunk = m.currentChunk + 1
				m.choices = m.allChunks[m.currentChunk]
				m.cursor = 0

			}

		case "enter", " ":

			if m.page == 0 {
				prevChoices[m.page] = m.namesCodeMap[selection]
				namesList, _, newMap := services.GetLeagueForCountry(m.namesCodeMap[selection])
				return newPageModel(namesList, newMap, 1, false), nil
			}
			if m.page == 1 {
				prevChoices[m.page] = m.namesCodeMap[selection]
				namesList, _, newMap := services.GetTeamsByLeagueId(m.namesCodeMap[selection])
				return newPageModel(namesList, newMap, 2, false), nil
			}
			if m.page == 2 {
				prevChoices[m.page] = m.namesCodeMap[selection]
				namesList, _, newMap := services.GetSquad(m.namesCodeMap[selection])
				return newPageModel(namesList, newMap, 3, false), nil
			}
			if m.page == 3 {
				prevChoices[m.page] = m.namesCodeMap[selection]
				statList, _, statMap := services.GetPlayerStats(m.namesCodeMap[selection])
				if len(statList) == 0 {
					statList = append(statList, "No stats for player")
				}
				return newPageModel(statList, statMap, 4, true), nil
			}

		case "b":
			newPage := (m.page - 1)
			lastChoice := prevChoices[newPage-1]

			if newPage == 0 {
				return initialModel(), nil
			}
			if newPage == 1 {
				namesList, _, newMap := services.GetLeagueForCountry(lastChoice)
				return newPageModel(namesList, newMap, 1, false), nil
			}
			if newPage == 2 {
				namesList, _, newMap := services.GetTeamsByLeagueId(lastChoice)
				return newPageModel(namesList, newMap, 2, false), nil
			}
			if newPage == 3 {
				namesList, _, newMap := services.GetSquad(lastChoice)
				return newPageModel(namesList, newMap, 3, false), nil
			}

		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	if m.isPlayerStat {
		s = "Player Stats:\n\n"
		for _, v := range m.allChoices {
			s += v
			s += "\n"

		}
		s += "\nPress q to quit or b to go back\n"

	} else {
		s = "Select from below: \n\n"

		for i, choice := range m.choices {

			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			choiceColor := aurora.Green(choice)
			if cursor == " " {
				s += fmt.Sprintf("%s  %s\n", aurora.Green(cursor), choice)
			} else {
				s += fmt.Sprintf("%s  %s\n", aurora.Green(cursor), choiceColor)
			}

		}
		if m.page == 0 {
			s += "\nPress q to quit\n"
		} else {
			s += "\nPress q to quit or b to go back\n"
		}

	}
	return s
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
