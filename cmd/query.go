package main

import (
	"flag"
	"fmt"
	"github.com/mikemrm/go-griddy"
)

func main() {
	config_file := flag.String("config", "", "Path to toml configuration file.")
	url := flag.String("url", "", "URL to be queried")
	meterid := flag.String("meterid", "", "Your meter ID")
	memberid := flag.String("memberid", "", "Your Member ID")
	settlement := flag.String("settlement", "", "Settlement Point (ex: LZ_HOUSTON)")
	flag.Parse()

	griddy := griddy.New()
	if *config_file != "" {
		err := griddy.LoadConfig(*config_file)
		if err != nil {
			panic(err)
		}
	}
	if *url != "" {
		griddy.Url = *url
	}
	if *meterid != "" {
		griddy.Meter = *meterid
	}
	if *memberid != "" {
		griddy.Member = *memberid
	}
	if *settlement != "" {
		griddy.Settlement = *settlement
	}
	response, err := griddy.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Time   : %s\n", response.Now.Date)
	fmt.Printf("Settlement    : %s\n", response.Now.Settlement)
	fmt.Printf("Price Type    : %s\n", response.Now.PriceType)
	fmt.Printf("Price         : %.2f %s\n", response.Now.Price, response.Now.PriceDisplaySign)
	fmt.Printf("Value Score   : %d\n", response.Now.ValueScore)
	fmt.Printf("Mean Price    : %.2f %s\n", response.Now.MeanPrice, response.Now.PriceDisplaySign)
	fmt.Printf("Diff Mean     : %.2f %s\n", response.Now.DiffMean, response.Now.PriceDisplaySign)
	fmt.Printf("High          : %.2f %s\n", response.Now.High, response.Now.PriceDisplaySign)
	fmt.Printf("Low           : %.2f %s\n", response.Now.Low, response.Now.PriceDisplaySign)
	fmt.Printf("Local Time    : %s\n", response.Now.DateLocalTZ)
	fmt.Printf("---------------\n")
	fmt.Printf("Data TTL      : %d Sec.\n", response.TTL)
}
