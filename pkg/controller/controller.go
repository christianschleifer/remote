package controller

import (
	"errors"
	"github.com/ChristianSchleifer/mremoteng/pkg/controller/api"
)

type controller struct {
	source  api.ConfigSource
	handler api.ConnectionHandler

	collection  api.Collection
	connections []*api.Connection
}

func NewController(source api.ConfigSource, handler api.ConnectionHandler) api.Controller {
	c := &controller{
		source:  source,
		handler: handler,
	}

	c.collection, c.connections = c.source.GetConfig()
	return c
}

func (c *controller) GetCollection() api.Collection {
	return c.collection
}

func (c *controller) ConnectionSelectedHandler(id uint) {
	connection, err := c.getConnectionById(id)
	if err != nil {
		panic(err)
	}

	c.handler.Handle(*connection)
}

func (c *controller) TransferControlForUIToHandler() bool {
	return c.handler.TransferControlForUI()
}

func (c *controller) getConnectionById(id uint) (*api.Connection, error) {
	for _, connection := range c.connections {
		if connection.Id == id {
			return connection, nil
		}
	}
	return nil, errors.New("no connection found")
}
