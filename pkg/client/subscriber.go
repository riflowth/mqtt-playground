package client

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client mqtt.Client
}

type MessageCombiner struct {
	seqs     int
	messages []string
}

type SensorData struct {
	NodeId       string
	Time         string
	Humidity     float64
	Temperature  float64
	ThermalArray string
}

var MessageMap map[string]MessageCombiner

// Create new mqtt subscriber instance
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

	MessageMap = make(map[string]MessageCombiner)

	return &Subscriber{client: client}, nil
}

// Subscribe to a topic
func (subscriber *Subscriber) Subscribe(topic string) error {
	if token := subscriber.client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Unsubscribe from a topic
func (subscriber *Subscriber) Unsubscribe(topic string) error {
	if token := subscriber.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Gets called everytime broker sends message to it
func onSubscribe(client mqtt.Client, message mqtt.Message) {
	// log data received
	log.Printf("recv: %s | topic: %s \n", message.Payload(), message.Topic())

	// convert JSON to string
	var payload Chunk
	json.Unmarshal(message.Payload(), &payload)

	//combine chunks of message
	combiner, found := MessageMap[payload.Id]
	if !found {
		initArray := []string{}
		for i := 0; i < payload.Seq; i++ {
			initArray = append(initArray, string(rune(i)))
		}
		MessageMap[payload.Id] = MessageCombiner{
			seqs:     payload.Seq,
			messages: initArray,
		}
	} else {
		combiner.messages[payload.Seq] = payload.Value
	}
	//insert to db
	if combiner.seqs-1 == payload.Seq {
		recv := strings.Split(strings.Join(combiner.messages, ""), " ")
		humidity, error := strconv.ParseFloat(recv[3], 64)
		if error != nil {
			panic(error)
		}
		temperature, error := strconv.ParseFloat(recv[4], 64)
		if error != nil {
			panic(error)
		}

		data := SensorData{
			NodeId:       recv[0],
			Time:         recv[1] + " " + recv[2],
			Humidity:     humidity,
			Temperature:  temperature,
			ThermalArray: recv[5],
		}
		log.Printf("Id:%v\n Time:%v\n Humidity:%v\n Temperature:%v\n ThermalArray:%v\n", data.NodeId, data.Time, data.Humidity, data.Temperature, data.ThermalArray)
	}
}
