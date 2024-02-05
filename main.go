package main

import (
	"crypto/rand"
	"devbook/src/config"
	"devbook/src/router"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}

func main() {
	config.Load()

	fmt.Println("Running API! Port: ", config.Port)
	fmt.Println(config.ConnectionString)
	fmt.Println(config.SecretKey)

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
