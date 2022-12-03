package auth

import (
	"utilsmod/config"
	"utilsmod/model"
)

type eventClient struct {
	conf config.Config
}

type EventClient interface {
	Subscribe(topic string) *model.ApiError
	Send(topic string, data []byte) *model.ApiError
}

func NewEventClient(conf config.Config) EventClient {
	return &eventClient{conf: conf}
}

func (e *eventClient) Subscribe(topic string) *model.ApiError {
	// Implement the subscription on a topic logic here
	return nil
}
func (e *eventClient) Send(topic string, data []byte) *model.ApiError {

	// Implement the kafka message sending logic here
	return nil
}
