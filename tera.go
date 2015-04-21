package main

import (
	"./api"
	"./config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	err := config.Load("tera.ini")
	if err != nil {
		log.Println(err)
		return
	}

	err = api.Listdir("")
	if err != nil {
		log.Println(err)
		return
	}
}
