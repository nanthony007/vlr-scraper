package scraping

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nanthony007/vlr-scraper/pkg"
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"log"
	"os"
	"strings"
)

func ScrapePlayerData(url string, mapName string) []models.PlayerStats {
	var rowCount int
	var data []models.PlayerStats

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get player stats
	c.OnHTML("table.wf-table-inset.mod-overview > tbody > tr", func(e *colly.HTMLElement) {
		// now handle rows
		var rowData [12]string
		rowCount += 1
		// these were HIDDEN javascript tables and its actually ALL the games!
		if rowCount <= 100 {
			// this is iterating each character
			metrics := strings.Split(e.Text, " ")
			for i, metric := range metrics {
				trimmedMetric := strings.TrimSpace(metric)
				cleanedTemp1 := strings.Replace(trimmedMetric, "%", "", 1)
				cleanedTemp2 := strings.Replace(cleanedTemp1, "/", "", 2)
				cleanedMetric := strings.Replace(cleanedTemp2, "+", "", 2)
				rowData[i] = cleanedMetric
			}
			playerInfo := models.NewPlayerStat(rowData)
			data = append(data, playerInfo)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
		fmt.Println(data)
		models.PlayersToFile(data, mapName)
	})

	c.Visit(url)
	return data
}

func FindMapPages(url string) {
	var gameIds [][]string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get map game ids and names
	c.OnHTML("div.vm-stats-gamesnav.noselect.mod-long > div.vm-stats-gamesnav-item.js-map-switch", func(e *colly.HTMLElement) {
		// maps
		if e.Attr("data-disabled") == "0" && e.Attr("data-game-id") != "all" {
			gameId := strings.TrimSpace(e.Attr("data-game-id"))
			extractedText1 := strings.ReplaceAll(e.Text, "\t", "")
			extractedText2 := strings.ReplaceAll(extractedText1, "\n", "")
			gameInfo := make([]string, 2)
			gameInfo[0] = gameId
			gameInfo[1] = strings.Trim(extractedText2, "123456")
			gameIds = append(gameIds, gameInfo)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
		file, err := os.Create("maps.csv")
		pkg.CheckErr(err)
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"GameID", "Map"})

		for _, row := range gameIds {
			err := writer.Write(row)
			pkg.CheckErr(err)
		}
	})

	c.Visit(url)
}