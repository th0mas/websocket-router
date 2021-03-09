package signal

import "encoding/json"

type signalType int

const (
	undef signalType = iota

	// Subscribe tells the handler we want to subscribe to another channel
	Subscribe

	// Unsubscribe tells the handler to remove the client from that channel
	Unsubscribe

	// Message tells the handler to pass the payload onto the router for another
	// handler function to worry about.
	// Most messages like this - the handler will look up the "path" and call
	// a registered
	Message
)

var toString = map[signalType]string{
	Subscribe:   "SUBSCRIBE",
	Unsubscribe: "UNSUBSRIBE",
	Message:     "MESSAGE",
}

var toID = map[string]signalType{
	"SUBSCRIBE":   Subscribe,
	"UNSUBSCRIBE": Unsubscribe,
	"MESSAGE":     Message,
}

// Signal represents an induvidial message of any type
type Signal struct {
	Cmd    signalType
	Path   string
	Sender string

	Message string
}

// FromBytes creates a signal from the raw WebSocket bytes
func FromBytes(payload []byte) Signal {
	var s Signal
	json.Unmarshal(payload, &s)

	return s
}
