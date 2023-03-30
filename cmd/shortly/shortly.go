package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ASA11599/shortly/internal/app"
)

func registerSignals() chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	return signals
}

func waitForApp(a app.App) <-chan bool {
	done := make(chan bool)
	go func(c chan<- bool) {
		a.Start()
		c <- true
	}(done)
	return done
}

func main() {
	sigs := registerSignals()
	var fra app.App = app.NewFiberRedisApp()
	defer fra.Stop()
	select {
	case <-waitForApp(fra):
		fmt.Println("App finished")
	case s := <-sigs:
		fmt.Println("Received signal:", s)
	}
}
