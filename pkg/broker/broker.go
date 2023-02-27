package broker

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
)

func NewMqttBroker() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	broker := mqtt.New(nil)
	broker.AddHook(new(auth.AllowHook), nil)
	broker.AddHook(new(LoggingHook), nil)
	

	tcp := listeners.NewTCP("broker", ":1883", nil)
	if error := broker.AddListener(tcp); error != nil {
		log.Fatalln(error)
	}

	if error := broker.Serve(); error != nil {
		log.Fatalln(error)
	}

	sig := <-sigs
	log.Printf("caught signal (%s), stopping...", sig)
	broker.Close()
}
