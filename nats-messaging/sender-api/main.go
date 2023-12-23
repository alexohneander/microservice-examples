package main

import (
	"log"
	"sender-api/nats"
)

func main() {
	log.Println("Listening for messages...")
	nats.Subscribe()
}
