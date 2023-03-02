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

func (publisher *Publisher) Publish(topic string, payload string) error {
	chunks := seperateByte(payload)
	for _, chunk := range chunks {
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

func onPublish(client mqtt.Client, message mqtt.Message) {
	log.Printf("send: %s | topic: %s\n", message.Payload(), message.Topic())
}

func seperateByte(str string) []Chunk {
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

	return chunks
}
