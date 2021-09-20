package models

import (
	"encoding/csv"
	"os"
	"strconv"
)

type PlayerStats struct {
	Name               string
	Team               string
	ACS                int64
	K                  int64
	D                  int64
	A                  int64
	PlusMinus          int64
	ADR                float64
	HeadShotPercent    int64
	FirstKill          int64
	FirstDeath         int64
	FirstKillDeathDiff int64
}

func NewPlayerStat(values [12]string) PlayerStats {
	// fix types
	playerInfo := PlayerStats{}
	playerInfo.Name = values[0]
	playerInfo.Team = values[1]
	// convert to int64
	if s, err := strconv.ParseInt(values[2], 10, 64); err == nil {
		playerInfo.ACS = s
	}
	if s, err := strconv.ParseInt(values[3], 10, 64); err == nil {
		playerInfo.K = s
	}
	if s, err := strconv.ParseInt(values[4], 10, 64); err == nil {
		playerInfo.D = s
	}
	if s, err := strconv.ParseInt(values[5], 10, 64); err == nil {
		playerInfo.A = s
	}
	if s, err := strconv.ParseInt(values[6], 10, 64); err == nil {
		playerInfo.PlusMinus = s
	}
	if s, err := strconv.ParseFloat(values[7], 64); err == nil {
		playerInfo.ADR = s
	}
	if s, err := strconv.ParseInt(values[8], 10, 64); err == nil {
		playerInfo.HeadShotPercent = s
	}
	if s, err := strconv.ParseInt(values[9], 10, 64); err == nil {
		playerInfo.FirstKill = s
	}
	if s, err := strconv.ParseInt(values[10], 10, 64); err == nil {
		playerInfo.FirstDeath = s
	}
	if s, err := strconv.ParseInt(values[11], 10, 64); err == nil {
		playerInfo.FirstKillDeathDiff = s
	}
	return playerInfo
}

func convertPlayerToStringArray(player PlayerStats) []string {
	// handle types
	array := make([]string, 12)
	array[0] = player.Name
	array[1] = player.Team
	// convert to float64
	array[2] = strconv.FormatInt(player.ACS, 10)
	array[3] = strconv.FormatInt(player.K, 10)
	array[4] = strconv.FormatInt(player.D, 10)
	array[5] = strconv.FormatInt(player.A, 10)
	array[6] = strconv.FormatInt(player.PlusMinus, 10)
	array[7] = strconv.FormatFloat(player.ADR, 'f', 1, 64)
	array[8] = strconv.FormatInt(player.HeadShotPercent, 10)
	array[9] = strconv.FormatInt(player.FirstKill, 10)
	array[10] = strconv.FormatInt(player.FirstDeath, 10)
	array[11] = strconv.FormatInt(player.FirstKillDeathDiff, 10)
	return array
}

func PlayersToFile(players []PlayerStats, fileprefix string) {
	fp := fileprefix + "_players.csv"
	file, err := os.Create(fp)
	CheckErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Name", "Team", "ACS", "K", "D", "A", "PlusMinus", "ADR", "HeadShotPercent", "FirstKill", "FirstDeath", "FirstKillDeathDiff"}
	writer.Write(headers)

	for _, player := range players {
		writablePlayer := convertPlayerToStringArray(player)
		err := writer.Write(writablePlayer)
		CheckErr(err)
	}
}
