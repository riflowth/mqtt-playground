package client

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type Publisher struct {
	client mqtt.Client
}

type Chunk struct {
	Id    string
	Seq   int
	Value string
}

// Create new mqtt publisher instance
func NewPublisher(id string, hostname string) (*Publisher, error) {
	opts := mqtt.
		NewClientOptions().
		AddBroker("tcp://" + hostname).
		SetClientID(id).
		SetDefaultPublishHandler(onPublish)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Publisher{client: client}, nil
}

// publish data to a topic to broker
func (publisher *Publisher) Publish(topic string, payload string) error {
	chunks := separateByte(payload)
	for _, chunk := range chunks {
		// convert data to JSON format
		data, error := json.Marshal(chunk)
		log.Println(string(data))
		if error != nil {
			return error
		}
		if token := publisher.client.Publish(topic, 0, false, data); token.Wait() && token.Error() != nil {
			return token.Error()
		}
	}
	return nil
}

// log everytime it publishes
func onPublish(client mqtt.Client, message mqtt.Message) {
	log.Printf("send: %s | topic: %s\n", message.Payload(), message.Topic())
}

// separate read data into chunks with 250 bytes each
func separateByte(str string) []Chunk {
	BytePerChunk := 100

	var chunks []Chunk
	var i, start, end int
	id := uuid.New().String()

	for i, start, end = 0, 0, BytePerChunk; end < len(str); i, start, end = i+1, end, end+BytePerChunk {
		chunks = append(chunks, Chunk{
			Id:    id,
			Seq:   i,
			Value: str[start:end],
		})
	}
	chunks = append(chunks, Chunk{
		Id:    id,
		Seq:   i,
		Value: str[start:],
	})
	chunks = append([]Chunk{{
		Id:    id,
		Seq:   i + 1,
		Value: "",
	}}, chunks...)

	return chunks
}
