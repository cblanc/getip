package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Ip struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func logError(err error) {
	fmt.Printf("%s", err)
	os.Exit(1)
}

func prettyPrint(ip Ip) {
	data := [][]string{
		[]string{"Country", ip.Country},
		[]string{"Country Code", ip.CountryCode},
		[]string{"Region", ip.Region},
		[]string{"Region Name", ip.RegionName},
		[]string{"City", ip.City},
		[]string{"Latitude", strconv.FormatFloat(ip.Lat, 'f', 6, 64)},
		[]string{"Longitude", strconv.FormatFloat(ip.Lon, 'f', 6, 64)},
		[]string{"Timezone", ip.Timezone},
		[]string{"ISP", ip.Isp},
		[]string{"Organisation", ip.Org},
		[]string{"Autonomous System Number & Name", ip.As},
		[]string{"Querying IP Address (You)", ip.Query},
	}

	table := tablewriter.NewWriter(os.Stdout)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func extractAddress(args []string) string {
	if len(args) < 2 {
		fmt.Printf("\nLookup up current IP Address\n")
		return ""
	} else {
		address := args[1]
		fmt.Printf("\nLooking up: %s\n", address)
		return address
	}
}

func main() {
	address := extractAddress(os.Args)

	response, err := http.Get("http://ip-api.com/json/" + address)
	if err != nil {
		logError(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logError(err)
	}

	var ip Ip
	jsonError := json.Unmarshal(contents, &ip)

	if jsonError != nil {
		logError(jsonError)
	}

	prettyPrint(ip)
}
