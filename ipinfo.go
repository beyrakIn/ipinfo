package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	green  = color.Green
	red    = color.Red
	yellow = color.Yellow
)

var (
	client     = &http.Client{}
	domain     = "https://ifconfig.co/?ip="
	numWorkers = 1
	verbose    = false
	interval   = 1.0
	done       = make(chan bool)
)

func main() {
	header()

	// implementation of command line arguments
	for i, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			help()
		} else if arg == "-w" || arg == "--workers" {
			// convert arg to int
			numWorkers = int(os.Args[i+1][0] - '0')
		} else if arg == "-v" || arg == "--verbose" {
			verbose = true
		} else if arg == "-i" || arg == "--interval" {
			// convert arg to int
			interval = float64(os.Args[i+1][0] - '0')
		}
	}

	work := make(chan string)

	go func() {
		select {
		case <-done:
			os.Exit(0)
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		yellow("\nCTRL+C detected!!!\nThanks for using this tool.\n")
		done <- true
	}()

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			t := strings.TrimSpace(s.Text())

			if err := checkInput(t); err != nil {
				red(err.Error())
				continue
			}
			work <- t
		}
		close(work)
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doWork(work, wg)
	}
	wg.Wait()

}

func doWork(work chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for ip := range work {
		url := domain + ip

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			red(err.Error())
		}

		req = req.WithContext(ctx)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Gecko/41.0 Firefox/41.0")
		response, err := client.Do(req)
		if err != nil {
			red(err.Error())
		}

		data, _ := io.ReadAll(response.Body)
		var d IpInfo
		if err := json.Unmarshal(data, &d); err != nil {
			red(err.Error())
		}

		if verbose == true {
			green(d.String())
		} else {
			green(d.Detail())
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func header() {
	text := "IF8gX19fX18gIF8gIF9fICBfICBfX19fICBfX19fIA0KfCB8fCAoKV8pfCB8fCAgXHwgfHwgPT09fC8gKCkgXA0KfF98fF98ICAgfF98fF98XF9ffHxfX3wgIFxfX19fLw0K"
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	base64.StdEncoding.Decode(base64Text, []byte(text))
	yellow(string(base64Text))
	red("Press CTRL+C to exit\n\n")
}

func help() {
	// print help
	fmt.Println("Usage: ipinfo [options]")
	fmt.Println("Options:")
	fmt.Println("  -h, --help\t\t\t\tShow this help message and exit")
	fmt.Println("  -v, --verbose\t\t\t\tShow full information")
	fmt.Println("  -w, --workers\t\t\t\tNumber of workers(default: 1)")
	fmt.Println("  -i, --interval\t\t\tInterval between requests(default: 0.4)")

	os.Exit(0)
}

func checkInput(t string) error {
	if strings.ToLower(t) == "exit" {
		red("Bye bye!!!")
		done <- true
	}

	regexS := `^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$`
	re := regexp.MustCompile(regexS)
	if !re.MatchString(t) {
		return fmt.Errorf(t + " is not valid ip address!!!")
	}
	return nil

}

type IpInfo struct {
	IP         string  `json:"ip"`
	IPDecimal  int     `json:"ip_decimal"`
	Country    string  `json:"country"`
	CountryIso string  `json:"country_iso"`
	CountryEu  bool    `json:"country_eu"`
	RegionName string  `json:"region_name"`
	RegionCode string  `json:"region_code"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TimeZone   string  `json:"time_zone"`
	Asn        string  `json:"asn"`
	AsnOrg     string  `json:"asn_org"`
	UserAgent  struct {
		Product  string `json:"product"`
		Version  string `json:"version"`
		RawValue string `json:"raw_value"`
	} `json:"user_agent"`
}

func (info IpInfo) String() string {
	return fmt.Sprintf(
		"IP: %s\nIP Decimal: %d\nCountry: %s\nCountry ISO: %s\nCountry EU: %t\nRegion Name: %s\nRegion Code: %s\nCity: %s\nLatitude: %f\nLongitude: %f\nTime Zone: %s\nASN: %s\nASN Org: %s\nUser Agent:\n  Product: %s\n  Version: %s\n  Raw Value: %s\n\n",
		info.IP, info.IPDecimal, info.Country, info.CountryIso, info.CountryEu, info.RegionName, info.RegionCode, info.City, info.Latitude, info.Longitude, info.TimeZone, info.Asn, info.AsnOrg, info.UserAgent.Product, info.UserAgent.Version, info.UserAgent.RawValue,
	)
}

func (info IpInfo) Detail() string {
	return fmt.Sprintf(
		"IP: %s\nCountry: %s\nASN: %s\nASN Org: %s\nTime Zone: %s\n\n",
		info.IP, info.Country, info.Asn, info.AsnOrg, info.TimeZone,
	)
}
