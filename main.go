package main

import (
	"log"
	"net/http"
	"sigmadoc/routers"
)

func main() {
	err := http.ListenAndServe(":3000", routers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
