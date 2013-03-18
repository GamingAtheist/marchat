package main

import "encoding/json"

type Message struct {
	Sender     string
	Text       string
	Encryption bool
}

func Decode(msg []byte) (msgStr string, err error) {
	foo := new
}
