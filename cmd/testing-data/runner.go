package main

import (
	"github.com/nanthony007/vlr-scraper/pkg/models"
	"github.com/nanthony007/vlr-scraper/pkg/scraping"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	matchData := scraping.ScrapeMatchOverview(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
	)
	eco_rounds := scraping.ScrapeEconomy(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"Split", "23773",
	)
	other_rounds := scraping.ScrapeGame(
		"https://www.vlr.gg/13247/vision-strikers-vs-nuturn-champions-tour-korea-stage-1-masters-gf",
		"Split", "23773",
	)
	// very messy
	rounds := models.NewAllRoundsData(eco_rounds, other_rounds)
	dirPath := filepath.Join("data", "match_"+"23773", "maps", "Split")
	os.MkdirAll(dirPath, 0777)
	fPath := filepath.Join(dirPath, "time_series.csv")
	models.RoundsToFile(rounds, fPath)
	ymlData, _ := yaml.Marshal(&matchData)
	dirPath = filepath.Join("data", "match_"+"23773")
	os.MkdirAll(dirPath, 0777)
	ymlPath := filepath.Join(dirPath, "series_info.yaml")
	err := ioutil.WriteFile(ymlPath, ymlData, 0644)
	models.CheckErr(err)
	ioutil.WriteFile(ymlPath, ymlData, 0644)
}
