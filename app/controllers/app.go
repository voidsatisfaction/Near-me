package controllers

import (
	"near_me_server/app/api"
	"near_me_server/app/factory"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Myplace() revel.Result {
	latitude := c.Params.Query.Get("lat")
	longitude := c.Params.Query.Get("lng")

	city, err := api.YahooGetLocation(latitude, longitude)
	if err != nil {
		return c.RenderError(err)
	}

	events := factory.Events{}

	connpassEventNums := 100
	connpassEvents, err := api.ConnpassGetEvents(city, connpassEventNums)
	if err != nil {
		return c.RenderError(err)
	}

	// add connpass events
	for _, connpassEvent := range connpassEvents {
		event := &factory.Event{}
		event.ConnpassAssign(&connpassEvent)
		if event.WillHold() {
			events.Data = append(events.Data, *event)
		}
	}

	doorkeeperEvents, err := api.DoorkeeperGetEvents(city, 0)
	if err != nil {
		return c.RenderError(err)
	}

	// add doorkeeper events
	for _, doorkeeperEvent := range doorkeeperEvents {
		event := &factory.Event{}
		event.DoorkeeperAssign(&doorkeeperEvent)
		if event.WillHold() {
			events.Data = append(events.Data, *event)
		}
	}
	events.Size = len(events.Data)

	return c.RenderJSON(events)
}