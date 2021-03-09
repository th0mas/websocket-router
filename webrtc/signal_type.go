package webrtc

import (
	"bytes"
	"encoding/json"
	"errors"
)

// SignalType represents the type of the WebRTC signal
type SignalType int

const (
	undef SignalType = iota
	// VideoOffer represents the VIDEO_OFFER type
	VideoOffer

	// VideoAnswer represents the VIDEO_ANSWER type
	VideoAnswer

	// NewIceCandidate represents the NEW_ICE_CANDIDATE type
	NewIceCandidate

	// DeviceLeave represents the DEVICE_LEAVE type
	DeviceLeave

	// DeviceJoin represents the DEVICE_JOIN type
	DeviceJoin
)

var toString = map[SignalType]string{
	VideoOffer:      "VIDEO_OFFER",
	VideoAnswer:     "VIDEO_ANSWER",
	NewIceCandidate: "NEW_ICE_CANDIDATE",
	DeviceLeave:     "DEVICE_LEAVE",
	DeviceJoin:      "DEVICE_JOIN",
}

var toID = map[string]SignalType{
	"VIDEO_OFFER":       VideoOffer,
	"VIDEO_ANSWER":      VideoAnswer,
	"NEW_ICE_CANDIDATE": NewIceCandidate,
	"DEVICE_LEAVE":      DeviceLeave,
	"DEVICE_JOIN":       DeviceJoin,
}

func (s SignalType) String() string {
	return toString[s]
}

// MarshalJSON marshals the enum as a quoted json string
func (s SignalType) MarshalJSON() ([]byte, error) {
	buff := bytes.NewBufferString("")
	buff.WriteString(toString[s])
	buff.WriteString(`"`)

	return buff.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *SignalType) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)

	if err != nil {
		return err
	}

	id := toID[str]

	if id == undef {
		return errors.New("signaltype: could not find type")
	}

	*s = id

	return nil
}
