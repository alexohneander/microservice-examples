package nats

import (
	"encoding/json"
	"log"
	"sender-api/model"

	"github.com/nats-io/nats.go"
)

var natsURL = "nats://nats:4222"

// var natsURL = "nats://localhost:4222"

func Subscribe() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("form", func(m *nats.Msg) {
		var form model.Form
		err := json.Unmarshal(m.Data, &form)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Form Received: %v", form)
	})

	select {}
}
