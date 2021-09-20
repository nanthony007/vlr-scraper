package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"github.com/nanthony007/vlr-scraper/pkg/utils"
	"log"
	"strings"
)

func ScrapeGame(url string, mapName string, gameID string) {

	// designed to VISIT the match page, then find various HTML elements and call accessory functions to scrape them
	// accessory functions should take HTML elements and either return data which can be added to global scope
	// OR
	// write data to file (preferable) which can then be contacted later
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get map info
	var teamData []map[string]string
	mapInfo := make(map[string]string)
	c.OnHTML("div.vm-stats-game[data-game-id='"+gameID+"'] div.vm-stats-game-header", func(e *colly.HTMLElement) {
		// TODO: modularize
		e.ForEach("div", func(i int, elem *colly.HTMLElement) {
			if elem.Attr("class") == "map" {
				// extract map info
				duration := elem.ChildText("div.map-duration")
				// messy
				name := strings.TrimSpace(strings.Split(strings.TrimSpace(elem.Text), "\n")[0])
				choiceClass := elem.ChildAttr("div span span", "class")
				var choice string
				if strings.Contains(choiceClass, "mod-1") {
					choice = "left"
				} else if strings.Contains(choiceClass, "mod-2") {
					choice = "right"
				} else {
					choice = "unknown" // should be error //
				}
				mapInfo["duration"] = duration
				mapInfo["name"] = name
				mapInfo["choice"] = choice
			} else if classText := elem.Attr("class"); classText == "team" || classText == "team mod-right" {
				teamInfo := make(map[string]string)
				// extract team info
				name := elem.ChildText("div div.team-name")
				defense := elem.ChildText("div span.mod-ct")
				attack := elem.ChildText("div span.mod-t")
				total := elem.ChildText("div.score")
				teamInfo["name"] = name
				teamInfo["defense"] = defense
				teamInfo["attack"] = attack
				teamInfo["total"] = total
				if teamInfo != nil {
					teamData = append(teamData, teamInfo)
				}
			} else {
				// raise error
			}
		})
	})

	// get player stats
	var playerData []models.PlayerStats
	c.OnHTML("div.vm-stats-game", func(e *colly.HTMLElement) {
		// these were HIDDEN (display:none) javascript tables
		// have to do this looping here bc colly passes each table one at a time
		table := utils.ExtractPlayerData(e, gameID)
		for _, row := range table {
			playerData = append(playerData, row)
		}
	})

	// get round results
	var rounds models.RoundResults
	var results [][]string
	c.OnHTML("div.vm-stats-game[data-game-id='"+gameID+"'] div div div.vlr-rounds div.vlr-rounds-row", func(e *colly.HTMLElement) {
		// long games have multiple rows so need top level loop
		e.ForEach("div.vlr-rounds-row-col", func(i int, elem *colly.HTMLElement) {
			// parse each column
			if i > 0 {
				// first col parse teams
				roundNum := elem.ChildText("div.rnd-num")
				// roundInt, _ := strconv.Atoi(roundNum)
				elem.ForEach("div.mod-win", func(a int, e2 *colly.HTMLElement) {
					// TODO: below
					// TODO: modularize
					// parse 4 options from first col
					// a == 0 and mod-t , a == 1 and mod-ct etc etc
					// get src link
					class := e2.Attr("class")
					classes := strings.Split(class, " ")
					winner := strings.Trim(classes[len(classes)-1], "mod-") // last one is valuable
					imgLink := e2.ChildAttr("img", "src")
					// parse link for valuable info
					linkParts := strings.Split(imgLink, "/")
					resultType := strings.Split(linkParts[len(linkParts)-1], ".")[0]
					results = append(results, []string{roundNum, winner, resultType})
				})
			}
		})
		rounds = models.NewRoundResults(results)
	})

	c.OnScraped(func(r *colly.Response) {
		// this is where we do a lot of data processing
		fmt.Println("Finished")
		// models.PlayersToFile(playerData, mapName)
		mapData := models.NewMapInfo(teamData, mapInfo)
		fmt.Println(mapData)
		fmt.Println(playerData)
		fmt.Println(rounds)
	})

	c.Visit(url)
}
