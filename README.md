# Vinderman

A library for making requests to Epic Games APIs--primarily targeting fortnite related endpoints.

Simple client creation example:
```go
package main

import (
	"gitlab.com/8h9x/vinderman"
	"gitlab.com/8h9x/vinderman/auth"
	"gitlab.com/8h9x/vinderman/consts"
	"log"
	"net/http"
)

func main() {
	deviceAuth := auth.PayloadDeviceAuth{
		AccountID: "ab825d5f61124e35b84ac31a62974609",
		DeviceID:  "f21a6d5c2e974e8991f071e35dcf98fc",
		Secret:    "2D509AFB52AC40568FEE7AC4AE9D92FA",
	}

	credentials, err := auth.Authenticate(consts.FortniteSwitchClientID, consts.FortniteSwitchClientSecret, deviceAuth, true)
	if err != nil {
		log.Println(err)
	}

	httpClient := &http.Client{}
	vinderClient, err := vinderman.NewClient(httpClient, credentials)
	if err != nil {
		log.Println("Failed to construct client", err)
	}

	log.Println(vinderClient)
}

```