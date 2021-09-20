package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"log"
	"strings"
)

func ScrapeEconomy(url string, mapName string, gameID string) models.EconomyRounds {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// get round econs
	var economyRounds models.EconomyRounds
	var ecoRounds [][]string
	c.OnHTML("div.vm-stats-game[data-game-id='"+gameID+"'] div table.wf-table-inset.mod-econ tbody tr", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, elem *colly.HTMLElement) {
			// TODO: this will need to implement the ct/t detection as well (maybe not?)
			var eco []string
			// skip teams for now
			if i > 0 {
				// get winner and two econ values
				class := elem.ChildAttr("div.mod-win", "class")
				classes := strings.Split(class, " ")
				winner := strings.Trim(classes[len(classes)-1], "mod-")
				eco = append(eco, winner)
				elem.ForEach("div.rnd-sq", func(j int, e2 *colly.HTMLElement) {
					cleanText := strings.TrimSpace(e2.Text)
					if cleanText != "" {
						eco = append(eco, cleanText)
					} else {
						eco = append(eco, "-")
					}
				})
				ecoRounds = append(ecoRounds, eco)
			}
		})
	})

	c.OnScraped(func(r *colly.Response) {
		// this is where we do a lot of data processing
		fmt.Println("Finished")
		var actualRounds [][]string
		for _, round := range ecoRounds {
			if len(round) == 3 {
				actualRounds = append(actualRounds, round)
			}
		}
		economyRounds = models.NewEconomyRounds(actualRounds)
	})

	fullUrl := url + "/?game=" + gameID + "&tab=economy"
	c.Visit(fullUrl)

	return economyRounds
}
