package main

import (
	"log"

	"github.com/zubairhamed/canopus"
)

func main() {
	conn, err := canopus.DialDTLS("localhost:5684", "canopus", "secretPSK")
	if err != nil {
		panic(err.Error())
	}

	req := canopus.NewRequest(canopus.MessageConfirmable, canopus.Get, canopus.GenerateMessageID())
	req.SetStringPayload("Hello, canopus")
	req.SetRequestURI("/hello")

	resp, err := conn.Send(req)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Got Response:" + resp.GetMessage().GetPayload().String())
}
