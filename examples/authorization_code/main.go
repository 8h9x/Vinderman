package main

import (
	"fmt"
	"gitlab.com/8h9x/vinderman"
	"gitlab.com/8h9x/vinderman/auth"
	"gitlab.com/8h9x/vinderman/consts"
	"log"
	"net/http"
)

func main() {
	httpClient := &http.Client{}

	var code string

	fmt.Printf("Enter an auth code from https://www.epicgames.com/id/api/redirect?clientId=%s&responseType=code:\n", consts.FortniteNewIOSClientID)
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatal(err)
	}

	authCodePayload := auth.PayloadAuthorizationCode{
		Code: code,
	}

	credentials, err := auth.Authenticate(httpClient, consts.FortniteNewIOSClientID, consts.FortniteNewIOSClientSecret, authCodePayload, true)
	if err != nil {
		log.Println(err)
	}

	_, err = vinderman.NewClient(httpClient, credentials)
	if err != nil {
		log.Println("Failed to construct client", err)
	}

	log.Println("Vinderman client successfully created")
}
