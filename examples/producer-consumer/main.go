package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	protocolbytes "github.com/mateusfdl/protocol-bytes"
	"golang.org/x/exp/rand"
)

func main() {
	messageChan := make(chan protocolbytes.Buffer, 10)
	go startConsumers(messageChan)
	go startProducers(messageChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}

func startConsumers(messageChan <-chan protocolbytes.Buffer) {
	for message := range messageChan {
		HandleMessage(message)
	}
}

func startProducers(messageChan chan<- protocolbytes.Buffer) {
	for {
		seed := rand.Intn(2)
		if seed == 0 {
			messageChan <- SendHelloMeassage()
		} else {
			messageChan <- SendThumbsUpMessage()
		}
		time.Sleep(450 * time.Millisecond)
	}
}
