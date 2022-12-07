package auth

import (
	"fmt"
	"time"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/logger"
	"github.com/arshabbir/utils/model"
)

type eventClient struct {
	conf *config.Config
	l    logger.Logger
}

type EventClient interface {
	Subscribe(topic string) *model.ApiError
	Send(topic string, data []byte) *model.ApiError
}

func NewEventClient(conf *config.Config, l logger.Logger) EventClient {
	return &eventClient{conf: conf, l: l}
}

func (e *eventClient) Subscribe(topic string) *model.ApiError {
	// Implement the subscription on a topic logic here
	e.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "EventClient", Message: fmt.Sprintf("Subscription request  : %s", topic)})
	return nil
}
func (e *eventClient) Send(topic string, data []byte) *model.ApiError {
	// Implement the kafka message sending logic here
	e.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "EventClient", Message: fmt.Sprintf("Send request, Topic   : %s, Data : %s", topic, data)})
	return nil
}
