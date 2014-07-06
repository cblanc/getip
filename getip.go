package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	address := ""

	if len(os.Args) < 2 {
		fmt.Printf("Lookup up current IP Address\n")
	} else {
		address = os.Args[1]
		fmt.Printf("Looking up: %s\n", address)
	}
	fmt.Printf("----------------------------\n")

	response, err := http.Get("http://ip-api.com/json/" + address)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		var ip Ip
		jsonError := json.Unmarshal(contents, &ip)
		if jsonError != nil {
			fmt.Printf("%s", jsonError)
			os.Exit(1)
		}
		fmt.Printf("Country: %s\n", ip.Country)
		fmt.Printf("Country Code: %s\n", ip.CountryCode)
		fmt.Printf("Region: %s\n", ip.Region)
		fmt.Printf("Region Name: %s\n", ip.RegionName)
		fmt.Printf("City: %s\n", ip.City)
		fmt.Printf("Latitude: %s\n", ip.Lat)
		fmt.Printf("Longitude: %s\n", ip.Lon)
		fmt.Printf("Timezone: %s\n", ip.Timezone)
		fmt.Printf("ISP: %s\n", ip.Isp)
		fmt.Printf("Organisation: %s\n", ip.Org)
		fmt.Printf("Autonomous System Number & Name: %s\n", ip.As)
		fmt.Printf("Querying IP Address (You): %s\n\n", ip.Query)
	}
}
