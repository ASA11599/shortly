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
	res := make(chan error)
	go func(c chan<- error) {
		c <- a.Start()
	}(res)
	return res
}

func main() {
	sigs := registerSignals()
	var sa app.App = app.GetShortlyApp()
	defer func() {
		if err := sa.Stop(); err != nil {
			fmt.Println("App stopped with error:", err)
		}
	}()
	select {
	case err := <-waitForApp(sa):
		fmt.Println("App finished with error:", err)
	case s := <-sigs:
		fmt.Println("Received signal:", s)
	}
}
