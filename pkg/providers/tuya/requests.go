package tuya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewTuya(deviceId string, deviceType string) Provider {
	if len(ClientId) == 0 || len(ClientSecret) == 0 {
		log.Fatal().Msg("Tuya environment variables not set")
	}

	return Provider{
		clientId:     ClientId,
		clientSecret: ClientSecret,
		deviceId:     deviceId,
		deviceType:   deviceType,
		accessToken:  "",
		currentTime:  "",
	}
}

func (t *Provider) AuthenticationRequest() error {
	client := http.Client{}
	authEndpoint := fmt.Sprintf("%s/%s", BaseDomain, AuthEndpoint)
	authRequest, err := http.NewRequest("GET", authEndpoint, nil)
	if err != nil {
		return err
	}

	t.currentTime = getCurrentTime()

	requestHeaders := map[string]string{
		"t":           t.currentTime,
		"sign":        t.generateAuthSignature(),
		"client_id":   t.clientId,
		"sign_method": "HMAC-SHA256",
	}

	setRequestHeaders(authRequest, requestHeaders)
	authResponse, err := sendRequest(&client, authRequest)

	var authResponseJson AuthResponse
	if err := json.Unmarshal(authResponse, &authResponseJson); err != nil {
		return err
	}

	t.accessToken = authResponseJson.Result.AccessToken
	return nil
}

func (t Provider) ToggleDeviceRequest(toggle bool) error {
	client := http.Client{}
	deviceEndpoint := fmt.Sprintf(
		"%s/%s/%s/%s", BaseDomain, DevicesEndpoint, t.deviceId, DevicesCommandsEndpoint,
	)

	requestBody := DeviceRequest{
		Commands: []DeviceRequestParameters{
			{
				Code:  t.deviceType,
				Value: toggle,
			},
		},
	}

	requestHeaders := map[string]string{
		"t":            t.currentTime,
		"sign":         t.generateRequestSignature(),
		"client_id":    t.clientId,
		"sign_method":  "HMAC-SHA256",
		"access_token": t.accessToken,
		"Content-Type": "application/json",
	}

	requestBodyJson, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	deviceRequest, err := http.NewRequest("POST", deviceEndpoint, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return err
	}

	setRequestHeaders(deviceRequest, requestHeaders)
	deviceResponse, err := sendRequest(&client, deviceRequest)
	if err != nil {
		return err
	}

	var deviceResponseJson DeviceResponse
	if err := json.Unmarshal(deviceResponse, &deviceResponseJson); err != nil {
		return err
	}

	return nil
}
