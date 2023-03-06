package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/riflowth/mqtt-lab/pkg/client"
	"github.com/riflowth/mqtt-lab/pkg/repositories"
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
			&cli.StringFlag{
				Name:  "topic",
				Usage: "To define topic",
			},
			&cli.StringFlag{
				Name:  "hostname",
				Usage: "To define hostname",
			},
			&cli.StringFlag{
				Name:  "influx-token",
				Usage: "To define influx token",
			},
			&cli.StringFlag{
				Name:  "influx-org",
				Usage: "To define influx org",
			},
			&cli.StringFlag{
				Name:  "influx-bucket",
				Usage: "To define influx bucket",
			},
			&cli.StringFlag{
				Name:  "influx-hostname",
				Usage: "To define influx hostname",
			},
		},
		// Define an action after execution this program
		Action: func(ctx *cli.Context) error {
			// Open channel to receive SIGINT or SIGTERM to terminate a process gracfully
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

			// Read `id` from execution flag for assigning subscriber id
			id := ctx.String("id")
			if id == "" {
				return errors.New("flag id is required, try --help for more information")
			}

			// Read `topic` from execution flag for subscribing to specific topic
			topic := ctx.String("topic")
			if topic == "" {
				return errors.New("flag topic is required, try --help for more information")
			}

			// Read host from execution flag for choosing hostname
			hostname := ctx.String("hostname")
			if hostname == "" {
				return errors.New("flag hostname is required, try --help for more information")
			}

			influxHostName := ctx.String("influx-hostname")
			if influxHostName == "" {
				log.Println("influx-hostname is not defined using default hostname `http://localhost:8086`")
				influxHostName = "http://localhost:8086"
			}

			token := ctx.String("influx-token")
			if token == "" {
				return errors.New("flag token is required, try --help for more information")
			}

			org := ctx.String("influx-org")
			if org == "" {
				log.Println("org is not defined using default org `admin`")
			}

			bucket := ctx.String("influx-bucket")
			if bucket == "" {
				log.Println("bucket is not defined using default bucket `mqtt`")
				bucket = "mqtt"
			}
			sensorRepository := repositories.NewSensorRepository(influxHostName, token, org, bucket)

			// Intialize subscriber with specific id from execution flag
			subscriber, error := client.NewSubscriber(id, hostname, sensorRepository)
			if error != nil {
				return error
			}

			// Start subscribe with specific flag from execution `topic` flag
			log.Printf("subscribing to topic %s", topic)
			subscriber.Subscribe(topic)

			// Wait for SIGINT or SIGTERM to terminate a subscriber process
			sig := <-sigs
			log.Printf("caught signal (%s), stopping...", sig)
			// Unsubscribe from specific topic
			subscriber.Unsubscribe(topic)

			return nil
		},
	}

	// Start application with input execution flags
	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}
