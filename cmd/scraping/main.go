package main

import (
	"encoding/csv"
	"github.com/nanthony007/vlr-scraper/pkg/scraping"
	"github.com/nanthony007/vlr-scraper/pkg/utils"
	"os"
)

func main() {

	url := "https://www.vlr.gg/34979/envy-vs-gambit-esports-valorant-champions-tour-stage-3-masters-berlin-gf"
	scraping.FindMapPages(url)

	file, err := os.Open("maps.csv")
	utils.CheckErr(err)
	defer file.Close()

	reader := csv.NewReader(file)

	maps, err := reader.ReadAll()
	utils.CheckErr(err)

	var urls []string
	for i, mapData := range maps {
		if i > 1 {
			newUrl := url + "/?game=" + mapData[0] + "&tab=overview"
			urls = append(urls, newUrl)
		}
	}
	scraping.ScrapePlayerData(urls, "BIND")
}
