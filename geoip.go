package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"path"

	"github.com/oschwald/geoip2-golang"
)

type Record struct {
	IP                string  `json:",omitempty"`
	City              string  `json:",omitempty"`
	Continent         string  `json:",omitempty"`
	Country           string  `json:",omitempty"`
	ISO               string  `json:",omitempty"`
	IsInEuropeanUnion bool    `json:""`
	Latitude          float64 `json:",omitempty"`
	Longitude         float64 `json:",omitempty"`
	PostalCode        string  `json:",omitempty"`
	TimeZone          string  `json:",omitempty"`
}

var lang string
var ipStr string

func main() {
	// Prepare input
	flag.StringVar(&lang, "lang", "en", "enable locale")
	flag.StringVar(&ipStr, "ip", "", "IP to resolve")
	flag.Parse()

	var raw *geoip2.City
	var record Record

	defer func() {
		record.IP = ipStr

		if recover() == nil {
			// fill by raw if no error
			record.City = raw.City.Names[lang]
			record.Continent = raw.Continent.Names[lang]
			record.Country = raw.Country.Names[lang]
			record.ISO = raw.Country.IsoCode
			record.IsInEuropeanUnion = raw.Country.IsInEuropeanUnion
			record.Latitude = raw.Location.Latitude
			record.Longitude = raw.Location.Longitude
			record.PostalCode = raw.Postal.Code
			record.TimeZone = raw.Location.TimeZone
		}

		// print to stdout
		js, err := json.Marshal(record)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(js))
	}()

	current, err := os.Executable()
	if err != nil {
		panic(err)
	}

	db, err := geoip2.Open(path.Join(path.Dir(current), "GeoLite2-City.mmdb"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipStr)
	if ip == nil {
		// wrong ip
		panic("Invalid IP")
	}
	raw, err = db.City(ip)
	if err != nil {
		panic(err)
	}

	return
}
