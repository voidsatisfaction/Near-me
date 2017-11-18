package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type DoorkeeperEvents []DoorkeeperEvent

type DoorkeeperEvent struct {
	Event struct {
		ID           uint32    `json:"id"`
		Title        string    `json:"title"`
		StartsAt     time.Time `json:"starts_at"`
		Lat          string    `json:"lat"`
		Long         string    `json:"long"`
		PublicURL    string    `json:"public_url"`
		Participants int       `json:"participants"`
		Waitlisted   int       `json:"waitlisted"`
	} `json:"event"`
}

func DoorkeeperGetEvents(city string, nums int, doorkeeperCH chan DoorkeeperEvents, errCH chan error) {
	key := os.Getenv("DOORKEEPER_KEY")
	host := "https://api.doorkeeper.jp/"
	url := host + "events?q=" + city + "&sort=starts_at"

	doorkeeperEvents := DoorkeeperEvents{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errCH <- err
		return
	}
	req.Header.Add("Authorization", "Bearer "+key)
	req.Header.Add("Content-Type", "application/json")
	// not parsed
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errCH <- err
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCH <- err
		return
	}
	defer resp.Body.Close()

	json.Unmarshal(body, &doorkeeperEvents)

	doorkeeperCH <- doorkeeperEvents
}
