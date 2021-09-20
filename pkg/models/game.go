package models

import (
	"encoding/csv"
	"os"
	"strconv"
)

type GameStats struct{}

// needs a better name
type Team struct {
	// stats for one game on a map
	Name    string
	Score   int
	Attack  int
	Defense int
}

func NewTeam(teamData map[string]string) Team {
	total, _ := strconv.Atoi(teamData["total"])
	atk, _ := strconv.Atoi(teamData["attack"])
	def, _ := strconv.Atoi(teamData["defense"])
	return Team{
		Name:    teamData["name"],
		Score:   total,
		Attack:  atk,
		Defense: def,
	}
}

type MapInfo struct {
	Name     string `yaml:"Name"`
	Duration string `yaml:"Duration"`
	Choice   int    `yaml:"Choice"`
	Team1    Team   `yaml:"Team1"`
	Team2    Team   `yaml:"Team2"`
}

func NewMapInfo(teamData []map[string]string, mapData map[string]string) MapInfo {
	if len(teamData) != 2 {
		// error
		panic("bad team val length")
	}
	return MapInfo{
		Name:     mapData["name"],
		Duration: mapData["duration"],
		Choice:   ParseChoice(mapData["choice"]),
		Team1:    NewTeam(teamData[0]),
		Team2:    NewTeam(teamData[1]),
	}
}

type RoundResult struct {
	Number int
	Winner string
	Kind   string
}

func NewRoundResult(num int, winner string, kind string) RoundResult {
	return RoundResult{
		Number: num,
		Winner: winner,
		Kind:   kind,
	}
}

type RoundResults struct {
	Rounds []RoundResult
}

func NewRoundResults(results [][]string) RoundResults {
	var resultsArray []RoundResult
	for _, round := range results {
		num, _ := strconv.Atoi(round[0])
		result := NewRoundResult(num, round[1], round[2])
		resultsArray = append(resultsArray, result)
	}
	return RoundResults{Rounds: resultsArray}
}

type EconomyRound struct {
	Number int
	Winner string
	Eco1   string
	Eco2   string
}

func NewEconomyRound(num int, winner string, eco1 string, eco2 string) EconomyRound {
	return EconomyRound{
		Number: num,
		Winner: winner,
		Eco1:   eco1,
		Eco2:   eco2,
	}
}

type EconomyRounds struct {
	Rounds []EconomyRound
}

func NewEconomyRounds(results [][]string) EconomyRounds {
	var resultsArray []EconomyRound
	for i, round := range results {
		result := NewEconomyRound(i+1, round[0], round[1], round[2])
		resultsArray = append(resultsArray, result)
	}
	return EconomyRounds{Rounds: resultsArray}
}

type RoundData struct {
	Number int
	Winner string
	Kind   string
	Eco1   string
	Eco2   string
}

func NewRoundData(ecoRound EconomyRound, resultRound RoundResult) RoundData {
	return RoundData{
		Number: resultRound.Number,
		Winner: resultRound.Winner,
		Kind:   resultRound.Kind,
		Eco1:   ecoRound.Eco1,
		Eco2:   ecoRound.Eco2,
	}
}

type AllRoundsData struct {
	Rounds []RoundData
}

func NewAllRoundsData(ecoRounds EconomyRounds, resultRounds RoundResults) AllRoundsData {
	var data []RoundData
	if len(ecoRounds.Rounds) == len(resultRounds.Rounds) {
		for i, _ := range resultRounds.Rounds {
			round := NewRoundData(ecoRounds.Rounds[i], resultRounds.Rounds[i])
			data = append(data, round)
		}
	}
	return AllRoundsData{Rounds: data}
}

func convertRoundToStringArray(round RoundData) []string {
	// handle types
	array := make([]string, 5)
	strNum := strconv.Itoa(round.Number)
	array[0] = strNum
	array[1] = round.Winner
	array[2] = round.Kind
	array[3] = round.Eco1
	array[4] = round.Eco1
	return array
}

func RoundsToFile(rounds AllRoundsData, fp string) {
	file, err := os.Create(fp)
	CheckErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"RoundNumber", "Winner", "ResultType", "Team1Economy", "Team2Economy"}
	writer.Write(headers)

	for _, round := range rounds.Rounds {
		writeableRound := convertRoundToStringArray(round)
		err := writer.Write(writeableRound)
		CheckErr(err)
	}
}
