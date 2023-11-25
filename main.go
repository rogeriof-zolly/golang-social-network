package main

import (
	"devbook/src/config"
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println("Running API! Port: ", config.Port)
	fmt.Println(config.ConnectionString)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
