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

func waitForApp(a app.App) <-chan error {
	done := make(chan error)
	go func(c chan<- error) {
		c <- a.Start()
	}(done)
	return done
}

func main() {
	sigs := registerSignals()
	var fra app.App = app.GetInstance()
	defer func() {
		if err := fra.Stop(); err != nil {
			fmt.Println("App stopped with error:", err)
		}
	}()
	select {
	case err := <-waitForApp(fra):
		fmt.Println("App finished with error:", err)
	case s := <-sigs:
		fmt.Println("Received signal:", s)
	}
}
