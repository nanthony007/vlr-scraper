package main

import (
	"fmt"
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"github.com/nanthony007/vlr-scraper/pkg/scraping"
)

func main() {
	eco_rounds := scraping.ScrapeEconomy(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"TEST", "23773",
	)
	other_rounds := scraping.ScrapeGame(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"TEST", "23773",
	)
	rounds := models.NewAllRoundsData(eco_rounds, other_rounds)
	fmt.Println(rounds)
}
