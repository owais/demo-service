package runner

import (
	"os"
	"os/signal"

	"github.com/owais/demo-service/pkg/listeners"
)

func Run(listeners ...listeners.Listener) {
	for _, l := range listeners {
		l.Start()
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
