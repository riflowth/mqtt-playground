package client

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client mqtt.Client
}

func NewSubscriber(id string) (*Subscriber, error) {
	opts := mqtt.
		NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID(id).
		SetDefaultPublishHandler(onSubscribe)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	if token := client.Subscribe("data/1", 0, nil); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Subscriber{client: client}, nil
}

func (subscriber *Subscriber) Subscribe(topic string) error {
	if token := subscriber.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (subscriber *Subscriber) Unsubscribe(topic string) error {
	if token := subscriber.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func onSubscribe(client mqtt.Client, message mqtt.Message) {
	log.Printf("recv: %s | topic: %s \n", message.Payload(), message.Topic())
}
