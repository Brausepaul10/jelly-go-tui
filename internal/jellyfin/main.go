package jellyfin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id string `json:"Id"`
}

type AuthDetails struct {
	AccessToken string `json:"AccessToken"`
	User        User   `json:"User"`
}

var initialPayload map[string]string = map[string]string{
	"Username": "Jannik",
	"Pw":       "jannik",
}

func GetUser() AuthDetails {
	client := &http.Client{}
	var auths AuthDetails
	marshalledBytes, _ := json.Marshal(initialPayload)
	bufferPayload := bytes.NewBuffer(marshalledBytes)
	req, err := http.NewRequest("POST", "https://jj.ofurth.de/Users/AuthenticateByName", bufferPayload)
	if err != nil {
		fmt.Println(err)
		return AuthDetails{}
	}
	req.Header.Set("Authorization", `MediaBrowser Client="GoTUI", Device="Linux", DeviceId="12345", Version="0.0.1"`)
	req.Header.Set("Content-type", "application/json")

	res, err2 := client.Do(req)

	if err2 != nil {
		fmt.Println(err2)
		return AuthDetails{}
	}
	defer res.Body.Close()

	fmt.Printf("Initial handshake response code: %d", res.StatusCode)

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&auths)

	return auths
}
