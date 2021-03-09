package main

import (
	s "github.com/th0mas/websocket-router/signal"
)

var route = map[string]func(sig *s.Signal) error{
	"":       func(sig *s.Signal) error { return nil },
	"webrtc": handleWebRTCfunc,
}

func handleWebRTCfunc(sig *s.Signal) error {

}
