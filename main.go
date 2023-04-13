package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	ht "github.com/imthaghost/merryGoRound/pkg/http"
)

func main() {
	// Configure a tor client
	tor := ht.Tor{
		MaxTimeout:         20 * time.Second,
		MaxIdleConnections: 10,
	}

	// new instance of tor client
	torClient := tor.New()

	// check your IP with AWS
	res, err := torClient.Get("https://checkip.amazonaws.com")
	if err != nil {
		log.Printf(fmt.Sprintf("%v", err))
	}

	log.Printf(fmt.Sprintf("body: [%v], [%v]", res, res.StatusCode))
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf(fmt.Sprintf(" [STATUS CODE: %d]", res.StatusCode))
	}

	ip := string(body)

	log.Printf("IP: %s", ip)

}
