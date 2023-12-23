package nats

import (
	"encoding/json"
	"form-api/model"
	"log"

	"github.com/nats-io/nats.go"
)

var natsURL = "nats://nats:4222"

// var natsURL = "nats://localhost:4222"

func Publish(form model.Form) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	formBytes, err := json.Marshal(form)
	if err != nil {
		log.Fatal(err)
	}

	nc.Publish("form", formBytes)
}
