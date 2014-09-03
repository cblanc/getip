package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"net/http"
	"os"
)

type Ip struct {
	Status      string
	Country     string
	CountryCode string
	Region      string
	RegionName  string
	City        string
	Zip         string
	Lat         string
	Lon         string
	Timezone    string
	Isp         string
	Org         string
	As          string
	Query       string
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
		[]string{"Latitude", ip.Lat},
		[]string{"Longitude", ip.Lon},
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
