package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	var data = []byte(`{From=19876543212&To=13216549878&Body=Test SMS from Restcomm&StatusCallback=http://status.callback.url}`)
	req, err := http.NewRequest("POST", "https://<accountSid>:<authToken>@cloud.restcomm.com/restcomm/2012-04-24/Accounts/<accountSid>/SMS/Messages", data)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
