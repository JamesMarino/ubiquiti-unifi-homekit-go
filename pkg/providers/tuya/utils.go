package tuya

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (t *Provider) ToggleDevice(toggle bool) {
	if err := t.AuthenticationRequest(); err != nil {
		log.Error().Stack().Err(err).Msg("error authenticating")
	}

	if err := t.ToggleDeviceRequest(toggle); err != nil {
		log.Error().Stack().Err(err).Msg("toggling device state")
	}
}

func (t *Provider) generateAuthSignature() string {
	clientTime := t.clientId + t.currentTime

	h := hmac.New(sha256.New, []byte(t.clientSecret))
	h.Write([]byte(clientTime))
	hash := hex.EncodeToString(h.Sum(nil))

	return strings.ToUpper(hash)
}

func (t *Provider) generateRequestSignature() string {
	clientTime := t.clientId + t.accessToken + t.currentTime

	h := hmac.New(sha256.New, []byte(t.clientSecret))
	h.Write([]byte(clientTime))
	hash := hex.EncodeToString(h.Sum(nil))

	return strings.ToUpper(hash)
}

func setRequestHeaders(httpRequest *http.Request, headers map[string]string) {
	for header, value := range headers {
		httpRequest.Header.Add(header, value)
	}
}

func sendRequest(client *http.Client, request *http.Request) ([]byte, error) {
	authResponse, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(authResponse.Body)
}

func getCurrentTime() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
