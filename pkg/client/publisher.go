package client

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	client mqtt.Client
}

func NewPublisher(id string) (*Publisher, error) {
	opts := mqtt.
		NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID(id).
		SetDefaultPublishHandler(onPublish)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Publisher{client: client}, nil
}

func (publisher *Publisher) Publish(topic string, payload any) error {
	if token := publisher.client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func onPublish(client mqtt.Client, message mqtt.Message) {
	log.Printf("send: %s | topic: %s\n", message.Payload(), message.Topic())
}
