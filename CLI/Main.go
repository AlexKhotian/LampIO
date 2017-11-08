package main

import (
	"LampIO/Client"
	"LampIO/Server"
	"flag"
	"log"
)

func main() {
	isClient := flag.Bool("client", false, "indicates if we are client")
	isServer := flag.Bool("server", false, "indicates if we are server")
	flag.Parse()
	log.Println(*isClient, *isServer)
	if *isClient && *isServer {
		log.Fatalln("Cannot creat client abd server at same time")
	}
	if *isClient {
		Client.StartClient()
	}
	if *isServer {
		Server.StartServer()
	}
}
