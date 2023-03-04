package broker

import (
	"bytes"

	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/packets"
)

type LoggingHook struct {
	mqtt.HookBase
}

func (hook *LoggingHook) ID() string {
	return "logging"
}

func (hook *LoggingHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnPublish,
		mqtt.OnPublished,
	}, []byte{b})
}

// This function run when MQTT client connect to this broker
func (hook *LoggingHook) OnConnect(client *mqtt.Client, packet packets.Packet) {
	hook.Log.Info().Str("client", client.ID).
		Msgf("client (id=%s | ip=%s) connected", client.ID, client.Net.Remote)
}

// This function run when MQTT client disconnect from this broker
func (hook *LoggingHook) OnDisconnect(client *mqtt.Client, error error, expire bool) {
	hook.Log.Info().Str("client", client.ID).
		Bool("expire", expire).
		Err(error).
		Msgf("client (id=%s | ip=%s) disconnected", client.ID, client.Net.Remote)
}

// Broker print out published message
func (hook *LoggingHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {

	hook.Log.Info().Str("client", cl.ID).
		// Message with topic and payload and id and ip form publisher
		Msgf("client (id=%s | ip=%s) published message with topic %s and payload %s", cl.ID, cl.Net.Remote, pk.TopicName, pk.Payload)

	return pk, nil

}
