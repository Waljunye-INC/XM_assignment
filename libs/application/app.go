package application

import (
	"OMS_assignment/libs/listeners"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

type app struct {
	listeners map[int]listeners.PortListener
}

func New(listeners map[int]listeners.PortListener) *app {
	return &app{
		listeners: listeners,
	}
}

func (app *app) Run(stopChan chan os.Signal) {
	for port, listener := range app.listeners {
		log.Info().Msg(fmt.Sprintf("running: %s. port: %v", listener.Info(), port))
		go func(listener listeners.PortListener, port int) {
			err := listener.Run(port)
			if err != nil {
				log.Error().Err(err)
				return
			}
		}(listener, port)
	}

	<-stopChan

	for _, listener := range app.listeners {
		err := listener.Stop()
		if err != nil {
			log.Error().Msg(fmt.Sprintf("stop listener: %v", err))
			return
		}
	}

	log.Info().Msg("all listeners stopped")
}
