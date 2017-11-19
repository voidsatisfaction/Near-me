package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type MyAddress struct {
	ResultSet struct {
		Address []string `json:"Address"`
	} `json:"ResultSet"`
}

func YahooGetLocation(latitude string, longitude string) ([]string, error) {
	// TODO move this data to conf
	var userLocation []string
	yahooAppID := os.Getenv("YAHOO_APP_KEY")
	resp, err := http.Get("https://map.yahooapis.jp/placeinfo/V1/get?appid=" + yahooAppID + "&lat=" + latitude + "&lon=" + longitude)
	if err != nil {
		return userLocation, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userLocation, err
	}

	var myAddress MyAddress
	json.Unmarshal(body, &myAddress)

	prefacture := myAddress.ResultSet.Address[0]
	village := myAddress.ResultSet.Address[1]
	// unicode japanese one character -> 3bytes
	city := prefacture[:len(prefacture)-3] + "市"
	if strings.Contains(village, city) {
		userLocation = append(userLocation, city)
	}
	userLocation = append(userLocation, prefacture)
	return userLocation, nil
}
