package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gbrlsnchs/jwt/v3"
)

type Payload struct {
	Embed string `json:"embed"`
	Sub   string `json:"sub"`
}

var hs = jwt.NewHS256([]byte("YOUR_PRIVATE_KEY"))

func main() {

	service := func(w http.ResponseWriter, r *http.Request) {
		payload := Payload{
			Embed: "YOUR_EMBED_ID",
			Sub:   "your.user@email.com",
		}

		token, _ := jwt.Sign(payload, hs)

		tmpl, err := template.ParseFiles("./frontend/index.html")

		if err != nil {
			log.Println(err)
		}

		tmpl.Execute(w, string(token))
	}

	http.HandleFunc("/", service)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
