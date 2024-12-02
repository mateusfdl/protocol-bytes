package main

import (
	"fmt"

	protocolbytes "github.com/mateusfdl/protocol-bytes"
)

type MessageHandler interface {
	Read(buf protocolbytes.Buffer) error
	Print()
}

type HelloMessage struct {
	name    string
	message string
}

type ThumbsUpMessage struct {
	name           string
	numberOfThumbs uint8
}

func HandleMessage(msg protocolbytes.Buffer) {
	id := msg.RUInt8()
	m := map[uint8]MessageHandler{
		0: &HelloMessage{},
		1: &ThumbsUpMessage{},
	}

	if handler, ok := m[id]; ok {
		err := handler.Read(msg)
		if err != nil {
			fmt.Println(err)
		}

		handler.Print()
	}
}

func (m *HelloMessage) Read(buf protocolbytes.Buffer) error {
	m.name = buf.RUTF()
	m.message = buf.RUTF()

	return nil
}

func (m *HelloMessage) Print() {
	fmt.Println("HelloMessage: ", m.name, m.message)
}

func (m *ThumbsUpMessage) Read(buf protocolbytes.Buffer) error {
	m.name = buf.RUTF()
	m.numberOfThumbs = buf.RUInt8()

	return nil
}

func (m *ThumbsUpMessage) Print() {
	fmt.Println("ThumbsUpMessage: ", m.name, m.numberOfThumbs)
}
