package main

import "github.com/lgln/canopus"

func main() {
	server := canopus.NewServer()
	server.ProxyOverHttp(true)

	server.ListenAndServe(":5683")
	<-make(chan struct{})

}
