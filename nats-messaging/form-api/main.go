package main

import (
	"form-api/model"
	"form-api/nats"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		var form model.Form

		form.Name = r.FormValue("name")
		form.Email = r.FormValue("email")
		form.Text = r.FormValue("text")

		log.Printf("Form Received: %v", form)
		nats.Publish(form)
	})

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
