package webrtc

// Signal contains all the fields in a WebRTC signal
type Signal struct {
	Type    SignalType
	Content string
	Sender  string
}
