package models

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseChoice(x string) int {
	if x == "right" {
		return 2
	} else if x == "left" {
		return 1
	} else {
		// raise error
		return -1
	}
}

func ExtractPlayerData(e *colly.HTMLElement, gameId string) (playerData []PlayerStats) {
	// TODO: add error return
	var data [12]string
	// only target game
	if e.Attr("data-game-id") == gameId {
		// skips the 'cumulative' table
		// might have to make this 'td' instead of 'tr' when we do agent support
		e.ForEach("div > div > table.wf-table-inset.mod-overview > tbody > tr", func(i int, elem *colly.HTMLElement) {
			metrics := strings.Split(elem.Text, " ")
			for i, metric := range metrics {
				trimmedMetric := strings.TrimSpace(metric)
				cleanedTemp1 := strings.Replace(trimmedMetric, "%", "", 1)
				cleanedTemp2 := strings.Replace(cleanedTemp1, "/", "", 2)
				cleanedMetric := strings.Replace(cleanedTemp2, "+", "", 2)
				data[i] = cleanedMetric
			}
			playerInfo := NewPlayerStat(data)
			playerData = append(playerData, playerInfo)
		})
	}
	return
}
