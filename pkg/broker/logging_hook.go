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

func (hook *LoggingHook) OnConnect(client *mqtt.Client, packet packets.Packet) {
	hook.Log.Info().Str("client", client.ID).
		Msgf("client (id=%s | ip=%s) connected", client.ID, client.Net.Remote)
}

func (hook *LoggingHook) OnDisconnect(client *mqtt.Client, error error, expire bool) {
	hook.Log.Info().Str("client", client.ID).
		Bool("expire", expire).
		Err(error).
		Msgf("client (id=%s | ip=%s) disconnected", client.ID, client.Net.Remote)
}
