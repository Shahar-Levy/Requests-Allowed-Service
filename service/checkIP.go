package service

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/RequestsAllowedService/model"
)

//RequestAllowed checks if a given IP belongs to a country in a list of given countries
func RequestAllowed(w http.ResponseWriter, r *http.Request, ip string, whitelist []string) bool {
	w.Write([]byte("Service is running"))

	countryIPMap := makeCountryIPMap()

	// return true if the ip is in the whitelisted country list
	for i := 0; i < len(whitelist); i++ {
		for _, v := range countryIPMap {
			if (v.CountryName == whitelist[i]) && countryHasIP(ip, v.Network) {
				return true
			}
		}
	}

	return false
}

//Creates a map of [geoname_id]=model.Country
func makeCountryIPMap() map[string]*model.Country {
	//read geoname_id, country_name from csv and creates map
	csvFileCountryLocations, _ := os.Open("GeoLite2-Country-Locations-en.csv")
	locationReader := csv.NewReader(bufio.NewReader(csvFileCountryLocations))
	countryIPMap := make(map[string]*model.Country)
	for {
		line, error := locationReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//checks if the country's id exists in the map
		if _, ok := countryIPMap[line[3]]; !ok {
			countryIPMap[line[0]] = &model.Country{CountryName: line[5]}
		}
	}
	//read networks from csv and append to the corresponding geoname_id
	csvFileCountryBlocks, _ := os.Open("GeoLite2-Country-Blocks-IPv4.csv")
	blocksReader := csv.NewReader(bufio.NewReader(csvFileCountryBlocks))
	for {
		line, error := blocksReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//only gives a network to country ids that have a country names
		if _, ok := countryIPMap[line[1]]; ok {
			countryIPMap[line[1]].Network = append(countryIPMap[line[1]].Network, line[0])
		}
	}
	return countryIPMap
}

func countryHasIP(ip string, network []string) bool {
	for _, i := range network {
		if i == ip {
			return true
		}
	}
	return false
}
