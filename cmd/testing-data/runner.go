package main

import "github.com/nanthony007/vlr-scraper/pkg/scraping"

func main() {
	scraping.ScrapeEconomy(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"TEST", "23773",
	)
	scraping.ScrapeGame(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"TEST", "23773",
	)
}
