package models

import (
	"github.com/nanthony007/vlr-scraper/pkg/utils"
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
	Name     string
	Duration string
	Choice   int
	Team1    Team
	Team2    Team
}

func NewMapInfo(teamData []map[string]string, mapData map[string]string) MapInfo {
	if len(teamData) != 2 {
		// error
		panic("bad team val length")
	}
	return MapInfo{
		Name:     mapData["name"],
		Duration: mapData["duration"],
		Choice:   utils.ParseChoice(mapData["choice"]),
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
	Results []RoundResult
}

func NewRoundResults(results [][]string) RoundResults {
	var resultsArray []RoundResult
	for _, round := range results {
		num, _ := strconv.Atoi(round[0])
		result := NewRoundResult(num, round[1], round[2])
		resultsArray = append(resultsArray, result)
	}
	return RoundResults{Results: resultsArray}
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
