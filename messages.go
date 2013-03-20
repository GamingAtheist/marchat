package main

import "bytes"
import "fmt"
import "encoding/json"
import "time"

const DateFormat = "2006-02-01 15:04:05"

type Message struct {
	Sender     string
	Text       []byte
	Encryption bool
	Control    bool
}

func DecodeMessage(msg []byte) (msgStr string, err error) {
	M := new(Message)
	err = json.Unmarshal(msg, &M)
	if err != nil {
		return
	}

	if M.Encryption {
		if !M.Control && len(config.Key) > 0 {
			var tmp []byte
			tmp, err = Decrypt(config.Key, M.Text)
			if err == nil {
				M.Text = tmp
			} else {
				M.Text = []byte(ShowError("[decryption error]"))
			}
			err = nil
		} else if !M.Control {
			M.Text = []byte(ShowError("[no secret key]"))
		}

		M.Text = []byte(fmt.Sprintf("%s %s", ShowSuccess("[encrypted]"),
			string(M.Text)))
	}

	if !M.Control {
		msgStr = fmt.Sprintf("<%s> %s: %s\n", time.Now().Format(DateFormat),
			M.Sender, string(M.Text))
	} else {
		msgStr = fmt.Sprintf("<%s> %s %s\n", time.Now().Format(DateFormat),
			M.Sender, string(M.Text))
		msgStr = ShowControl(msgStr)
	}
	return
}

func EncodeMessage(msg []byte, control bool) (wire []byte, err error) {
	msg = bytes.TrimSpace(msg)
	M := new(Message)
	if !control && len(config.Key) != 0 {
		msg, err = Encrypt(config.Key, msg)
		if err != nil {
			return
		}
		M.Encryption = true
	}
	M.Sender = config.User
	M.Text = msg
	M.Control = control
	wire, err = json.Marshal(&M)
	return
}
