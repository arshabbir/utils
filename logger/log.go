package logger

import (
	"log"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/model"
)

type logger struct {
	conf *config.Config
}

type Logger interface {
	Log(model.LogRequest)
}

func NewLogger(conf *config.Config) Logger {
	return &logger{conf: conf}
}

func (l *logger) Log(m model.LogRequest) {
	// implement logging logic based on conf.Loglevel
	log.Printf("\n%s : %s :  %s", m.Timestamp, m.ServiceName, m.Message)

}
