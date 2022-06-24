package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"regexp"
)

var (
	green  = color.Green
	red    = color.Red
	yellow = color.Yellow
)

func main() {
	header()

	ip := flag.String("ip", "", "Ex: ipinfo -ip 8.8.8.8")
	flag.Parse()

	if *ip == "" {
		green("Please run ipinfo -h")
		return
	}

	if IsValidIp(*ip) {
		url := makeUrl(*ip)
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		req = req.WithContext(ctx)
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		data, _ := io.ReadAll(response.Body)
		var d IpInfo
		if err := json.Unmarshal(data, &d); err != nil {
			red(err.Error())
		}

		green(fmt.Sprintf(`
	IP: %s
	Hostname: %s
	City: %s
	Region: %s
	Country: %s
	Coordinates: %s
	ASN: %s
	Postal: %s 
	Timezone: %s  `, d.IP, d.Hostname, d.City, d.Region, d.Country, d.Loc, d.Org, d.Postal, d.Timezone))

	}

}

func IsValidIp(ip string) bool {
	regexS := `^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$`
	re := regexp.MustCompile(regexS)
	if !re.MatchString(ip) {
		red(ip + " is not valid ip address!!!")
		return false
	}
	return true
}

func makeUrl(ip string) string {
	domain := "https://ipinfo.io/"
	url := domain + ip
	return url
}

func header() {
	head := `.___       .___        _____       
|   |_____ |   | _____/ ____\____  
|   \____ \|   |/    \   __\/  _ \ 
|   |  |_> >   |   |  \  | (  <_> )
|___|   __/|___|___|  /__|  \____/ 
    |__|            \/             `
	yellow(head + "\n")
}

type IpInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}
