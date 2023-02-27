package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/riflowth/mqtt-lab/pkg/client"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "MQTT publisher",
		Usage: "CLI to start MQTT publisher",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "id",
				Usage: "To define publisher id",
			},
		},
		Action: func(ctx *cli.Context) error {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			id := ctx.String("id")
			if id == "" {
				return errors.New("flag id is required, try --help for more information")
			}

			publisher, error := client.NewPublisher(id)
			if error != nil {
				return error
			}

			publisher.Publish("data/1", "test")

			sig := <-sigs
			log.Printf("caught signal (%s), stopping...", sig)

			return nil
		},
	}

	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}
