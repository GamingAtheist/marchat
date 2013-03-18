package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var config struct {
	User string
	Port string
	Key  []byte
}

var (
	Incoming = make(chan []byte, 0)
	Outgoing = make(chan []byte, 0)
)

const chatPort = 4001

func transmitterHandler(ws *websocket.Conn) { ws.Close() }

func receiverHandler(ws *websocket.Conn) {
	messages := make([][]byte, 0)
	msgCount := len(Outgoing)
	for i := 0; i < msgCount; i++ {
		messages = append(messages, <-Incoming)
	}

	wire, err := json.Marshal(messages)
	if err != nil {
		ws.Close()
	}
	ws.Write(wire)
}

func renameHandler(ws *websocket.Conn) {
	buf := new(bytes.Buffer)
	io.Copy(buf, ws)
	config.User = string(buf.Bytes())
}

func main() {
	fKeyFile := flag.String("k", "", "key file")
	fPort := flag.Int("p", 4000, "listening port")
	fUser := flag.String("u", "anonymous", "user to broadcast as")
	flag.Parse()

	config.Port = fmt.Sprintf("%d", *fPort)
	config.User = *fUser

	if *fKeyFile != "" {
		var err error
		config.Key, err = ReadKeyFromFile(*fKeyFile)
		if err != nil {
			log.Fatalf("[!] failed to load %s: %s\n", *fKeyFile,
				err.Error())
		}
	}

	http.HandleFunc("/", rootHandler)
	http.Handle("/socket", websocket.Handler(transmitterHandler))
	http.Handle("/incoming", websocket.Handler(receiverHandler))
	http.Handle("/name", websocket.Handler(renameHandler))
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func networkChat() {
	gaddr := selectInterface()
}

func transmit() {
}
