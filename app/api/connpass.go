package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ConnpassEvents struct {
	Data []ConnpassEvent `json:"events"`
}

func (ce *ConnpassEvents) getEventNum() int {
	return len(ce.Data)
}

type ConnpassEvent struct {
	EventID   uint32    `json:"event_id"`
	EventURL  string    `json:"event_url"`
	StartedAt time.Time `json:"started_at"`
	Lat       string    `json:"lat"`
	Lon       string    `json:"lon"`
	Title     string    `json:"title"`
	Address   string    `json:"address"`
	Place     string    `json:"place"`
}

func ConnpassGetEvents(userLocation []string, nums int, connpassCH chan []ConnpassEvent, errCH chan error) {
	if nums == 0 {
		nums = 10
	}
	var connpassEvents ConnpassEvents

	host := "https://connpass.com/api/v1/event/?"
	keyword := strings.Join(userLocation, ",")
	resp, err := http.Get(host + "keyword_or=" + keyword + "&count=" + strconv.Itoa(nums))
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

	json.Unmarshal(body, &connpassEvents)
	connpassCH <- connpassEvents.Data
}
