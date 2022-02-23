package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)
// Avenger represents a single hero
// type Avenger struct {
// 	RealName string `json:"real_name"`
// 	HeroName string `json:"hero_name"`
// 	Planet   string `json:"planet"`
// 	Alive    bool   `json:"alive"`
// }

// func (a *Avenger) isAlive() {
// 	a.Alive = true
// }

func main() {
	// avengers := []Avenger{
	// 	{
	// 		RealName: "Dr. Bruce Banner",
	// 		HeroName: "Hulk",
	// 		Planet:   "Midgard",
	// 	},
	// 	{
	// 		RealName: "Tony Stark",
	// 		HeroName: "Iron Man",
	// 		Planet:   "Midgard",
	// 	},
	// 	{
	// 		RealName: "Thor Odinson",
	// 		HeroName: "Thor",
	// 		Planet:   "Midgard",
	// 	},
	// }

	// avengers[1].isAlive()

	// jsonBytes, err := json.Marshal(avengers)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(jsonBytes))

	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}