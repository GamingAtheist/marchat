package main

import (
	"encoding/json"
)

type Message struct {
	Sender    string
	Message   []byte
	Encrypted bool
}

func (m *Message) Encode() (mj []byte, err error) {
	mj = make([]byte, 0)
	mj, err = json.Marshal(&m)
	return
}

func DecodeMessage(mj []byte) (m *Message, err error) {
	m = new(Message)
	err = json.Unmarshal(mj, &m)
	return
}
