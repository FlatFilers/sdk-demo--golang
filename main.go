package main

import (
	"log"

	"github.com/gbrlsnchs/jwt/v3"
)

/*
import jwt from 'jsonwebtoken'
  const token = jwt.sign({
    embed: '',
    sub: ''
  }, 'YOUR_FLATFILE_PRIVATE_KEY')
*/

type Payload struct {
	Embed string `json:"embed"`
	Sub string	`json:"sub"`
}

var hs = jwt.NewHS256([]byte("secret"))

func main() {
	payload := Payload{
		Embed: "a15b45c3-4743-4c16-851d-b93e1d1c8836",
		Sub: "maire@flatfile.io",
	}

	token, err := jwt.Sign(payload, hs)
	if err != nil {
		log.Fatal("error signing jwt")
	}

	log.Println(token)
}
