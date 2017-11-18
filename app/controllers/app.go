package controllers

import (
	"near_me_server/app/api"
	"near_me_server/app/factory"
	"near_me_server/app/services"

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

	es, err := services.GetAllEvents(city)
	if err != nil {
		return c.RenderError(err)
	}

	connpassEvents := es.Connpass
	doorkeeperEvents := es.Doorkeeper

	// add connpass events
	for _, connpassEvent := range connpassEvents {
		event := &factory.Event{}
		event.ConnpassAssign(&connpassEvent)
		if event.WillHold() {
			events.Data = append(events.Data, *event)
		}
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
