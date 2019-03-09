body := strings.NewReader(`From=19876543212&To=13216549878&Body=Test SMS from Restcomm&StatusCallback=http://status.callback.url`)
req, err := http.NewRequest("POST", "https://<accountSid>:<authToken>@cloud.restcomm.com/restcomm/2012-04-24/Accounts/<accountSid>/SMS/Messages", body)
if err != nil {
	// handle err
}
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

resp, err := http.DefaultClient.Do(req)
if err != nil {
	// handle err
}
defer resp.Body.Close()
