package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/riflowth/mqtt-lab/pkg/client"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "MQTT subscriber",
		Usage: "CLI to start MQTT subscriber",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "id",
				Usage: "To define subscriber id",
			},
		},
		Action: func(ctx *cli.Context) error {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			id := ctx.String("id")
			if id == "" {
				return errors.New("flag id is required, try --help for more information")
			}

			subscriber, error := client.NewSubscriber(id)
			if error != nil {
				return error
			}

			log.Println("subscribe to data/1")
			subscriber.Subscribe("data/1")

			sig := <-sigs
			log.Printf("caught signal (%s), stopping...", sig)
			subscriber.Unsubscribe("data/1")

			return nil
		},
	}

	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}
