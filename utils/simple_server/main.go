package main

import "log"

func main() {
	server := SimpleServer{}
	server.Init()
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
