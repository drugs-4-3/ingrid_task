package main

import (
	_ "github.com/joho/godotenv/autoload"
	myHttp "github.com/drugs-4-3/ingrid_task/http"
	"log"
)

func main() {
	srvr, err := myHttp.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server, listening on %s", srvr.Addr)
	if err := srvr.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}