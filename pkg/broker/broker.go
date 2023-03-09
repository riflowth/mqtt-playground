package broker

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
)

func NewMqttBroker(port string) {
	// Open channel to receive SIGINT or SIGTERM to terminate a process gracfully
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Intialize mochi-co broker
	broker := mqtt.New(nil)
	// Add mochi-co broker auth hook to allow any connections without auth
	broker.AddHook(new(auth.AllowHook), nil)
	// Add mochi-co broker custom logging hook to log every behavior we want
	broker.AddHook(new(LoggingHook), nil)

	// Create TCP listener to bind to mochi-co broker with port that user defined (default: 1883)
	tcp := listeners.NewTCP("broker", fmt.Sprintf(":%s", port), nil)
	if error := broker.AddListener(tcp); error != nil {
		log.Fatalln(error)
	}

	// Serve broker server, if an error exists then terminate the broker process
	if error := broker.Serve(); error != nil {
		log.Fatalln(error)
	}

	// Wait for SIGINT or SIGTERM to terminate a broker process
	sig := <-sigs
	log.Printf("caught signal (%s), stopping...", sig)
	// Stop broker
	broker.Close()
}
