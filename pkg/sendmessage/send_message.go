package sendmessage

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/credentials"
)

func SendMessage(w WhatsApp) {
	// endpoint API WhatsApp
	const endpoint = "https://api.zenvia.com/v2/channels/whatsapp/messages"
	// Create a new request using http
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}

	gc := credentials.GetCredentials()
	w.From = gc.From

	// Add authorization header to the req
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", gc.X_API_TOKEN)

	// Marshal the struct to JSON
	jsonReq, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}

	// Set the body of the request to the JSON
	req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonReq))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	// Read the response body
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
}
