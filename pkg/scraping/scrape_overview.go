package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nanthony007/vlr-scraper/pkg/utils"
	"log"
	"strconv"
	"strings"
)

func ScrapeMatchOverview(url string, mapName string) {
	// var rowCount int
	// var data []models.PlayerStats

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get teams
	var teams []string
	c.OnHTML("div.match-header-vs > a > div > div.wf-title-med", func(e *colly.HTMLElement) {
		teams = append(teams, strings.TrimSpace(e.Text))
	})

	// get match score and type
	var scores []int
	c.OnHTML("div.match-header-vs-score > div.js-spoiler > span", func(e *colly.HTMLElement) {
		textVal := strings.TrimSpace(e.Text)
		if textVal != ":" {
			val, err := strconv.Atoi(textVal)
			utils.CheckErr(err)
			scores = append(scores, val)
		}
	})

	// get map game ids and names
	var gameIds []int
	c.OnHTML("div.vm-stats-gamesnav.noselect.mod-long > div.vm-stats-gamesnav-item.js-map-switch", func(e *colly.HTMLElement) {
		// maps
		if e.Attr("data-disabled") == "0" && e.Attr("data-game-id") != "all" {
			gameId := strings.TrimSpace(e.Attr("data-game-id"))
			val, err := strconv.Atoi(gameId)
			utils.CheckErr(err)
			gameIds = append(gameIds, val)
			// visit those maps
			// c.Visit(<new urls>)
		}
	})

	// this was player scraping, but that will not be done here
	//// get player stats
	//c.OnHTML("table.wf-table-inset.mod-overview > tbody > tr", func(e *colly.HTMLElement) {
	//	// now handle rows
	//	var rowData [12]string
	//	rowCount += 1
	//	// these were HIDDEN javascript tables and its actually ALL the games!
	//	// BUT this includes the total for the match... i.e. all games... must filter that out
	//	if rowCount <= 100 {
	//		// this is iterating each character
	//		metrics := strings.Split(e.Text, " ")
	//		for i, metric := range metrics {
	//			trimmedMetric := strings.TrimSpace(metric)
	//			cleanedTemp1 := strings.Replace(trimmedMetric, "%", "", 1)
	//			cleanedTemp2 := strings.Replace(cleanedTemp1, "/", "", 2)
	//			cleanedMetric := strings.Replace(cleanedTemp2, "+", "", 2)
	//			rowData[i] = cleanedMetric
	//		}
	//		playerInfo := models.NewPlayerStat(rowData)
	//		data = append(data, playerInfo)
	//	}
	//})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
		//fmt.Println(data)
		//models.PlayersToFile(data, mapName)
		// TODO: modularize below
		matchResults := map[string]int{}
		matchResults[teams[0]] = scores[0]
		matchResults[teams[1]] = scores[1]
		fmt.Println(matchResults)
	})

	c.Visit(url)
}
