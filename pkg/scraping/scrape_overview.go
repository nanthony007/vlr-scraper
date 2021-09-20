package scraping

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"log"
	"strconv"
	"strings"
)

func ScrapeMatchOverview(url string) Series {
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
			models.CheckErr(err)
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
			models.CheckErr(err)
			gameIds = append(gameIds, val)
			// visit those maps
			// c.Visit(<new urls>)
		}
	})

	series := Series{}
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
		//fmt.Println(data)
		//models.PlayersToFile(data, mapName)
		// TODO: modularize below
		series.Team1.Name = teams[0]
		series.Team2.Name = teams[1]
		series.Team1.Score = scores[0]
		series.Team2.Score = scores[1]
	})

	c.Visit(url)
	return series
}

type Team struct {
	Name  string `yaml:"Name"`
	Score int    `yaml:"Score"`
}

type Series struct {
	Team1 Team `yaml:"Team1"`
	Team2 Team `yaml:"Team2"`
}
