package services

import (
	"log"
	"near_me_server/app/api"
)

type Events struct {
	Connpass   []api.ConnpassEvent
	Doorkeeper api.DoorkeeperEvents
}

func GetAllEvents(userLocation []string) (*Events, error) {
	// this is the number of event type
	typeNum := 2

	connpassCH := make(chan []api.ConnpassEvent)
	errCH := make(chan error)
	connpassEventNums := 100
	go api.ConnpassGetEvents(userLocation, connpassEventNums, connpassCH, errCH)

	doorkeeperCH := make(chan api.DoorkeeperEvents)
	go api.DoorkeeperGetEvents(userLocation, 0, doorkeeperCH, errCH)

	var connpassEvents []api.ConnpassEvent
	var doorkeeperEvents api.DoorkeeperEvents
	for i := 0; i < typeNum; i++ {
		select {
		case err := <-errCH:
			if err != nil {
				log.Fatal("Fetch events error!!")
				log.Fatal(err)
				return nil, err
			}
		case connpassEvents = <-connpassCH:
		case doorkeeperEvents = <-doorkeeperCH:
		}
	}
	es := &Events{
		Connpass:   connpassEvents,
		Doorkeeper: doorkeeperEvents,
	}
	return es, nil
}
