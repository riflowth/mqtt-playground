package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/riflowth/mqtt-lab/pkg/client"
	"github.com/riflowth/mqtt-lab/pkg/client/sensors"
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
			&cli.StringFlag{
				Name:  "topic",
				Usage: "To define topic",
			},
			&cli.StringFlag{
				Name:  "hostname",
				Usage: "To define hostname",
			},
		},
		// Define an action after execution this program
		Action: func(ctx *cli.Context) error {
			// Open channel to receive SIGINT or SIGTERM to terminate a process gracfully
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			// Read `id` from execution flag for assigning publisher id
			id := ctx.String("id")
			if id == "" {
				return errors.New("flag id is required, try --help for more information")
			}

			// Read `topic` from execution flag for publishing to specific topic
			topic := ctx.String("topic")
			if topic == "" {
				return errors.New("flag topic is required, try --help for more information")
			}

			// Read host from execution flag for choosing hostname
			hostname := ctx.String("hostname")
			if hostname == "" {
				return errors.New("flag hostname is required, try --help for more information")
			}

			// Initialize publisher with specific id from execution flag
			publisher, error := client.NewPublisher(id, hostname)
			if error != nil {
				return error
			}

			go ReadSensor(id, *publisher, topic)

			// Wait for SIGINT or SIGTERM to terminate a publisher process
			sig := <-sigs
			log.Printf("caught signal (%s), stopping...", sig)

			return nil
		},
	}

	// Start application with input execution flags
	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}

// Read sensors data and row amount of data
func ReadSensor(id string, publisher client.Publisher, topic string) {

	s := sensors.NewSensors()
	rows := sensors.GetNumRows()

	for i := 1; i < rows; i++ {
		d := sensors.Read(s)
		d = id + " " + d
		publisher.Publish(topic, d)
		time.Sleep(3 * time.Minute)
	}
}
