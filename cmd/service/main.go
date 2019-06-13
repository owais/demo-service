package main

import (
	"github.com/owais/demo-service/pkg/listeners/red"
	"github.com/owais/demo-service/pkg/runner"

	"github.com/owais/demo-contrib/listeners/blue"
)

func main() {
	runner.Run(red.Listener{}, blue.Listener{})
}
