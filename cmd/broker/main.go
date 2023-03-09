package main

import (
	"log"
	"os"

	"github.com/riflowth/mqtt-lab/pkg/broker"
	"github.com/urfave/cli"
)

func main() {
	// Initialize broker instance and start broker

	app := &cli.App{
		Name:  "MQTT broker",
		Usage: "CLI to start MQTT broker",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Usage: "To define port to exposed",
			},
		},
		// Define an action after execution this program
		Action: func(ctx *cli.Context) error {

			port := ctx.String("port")
			if port == "" {
				log.Println("flag port is not defined, using default port (1883) instead")
				port = "1883"
			}

			broker.NewMqttBroker(port)
			return nil
		},
	}

	// Start application with input execution flags
	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}
