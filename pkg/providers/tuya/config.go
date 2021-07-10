package tuya

import "os"

const (
	BaseDomain              string = "https://openapi.tuyaeu.com/v1.0"
	AuthEndpoint            string = "token?grant_type=1"
	DevicesEndpoint         string = "devices"
	DevicesCommandsEndpoint string = "commands"
)

var (
	ClientId     = os.Getenv("TUYA_CLIENT_ID")
	ClientSecret = os.Getenv("TUYA_CLIENT_SECRET")
)

type AuthResponse struct {
	Result AuthResult `json:"result"`
}

type AuthResult struct {
	AccessToken string `json:"access_token"`
}

type DeviceRequest struct {
	Commands []DeviceRequestParameters `json:"commands"`
}

type DeviceRequestParameters struct {
	Code  string `json:"code"`
	Value bool   `json:"value"`
}

type DeviceResponse interface {
}

type Provider struct {
	clientId     string
	clientSecret string
	accessToken  string
	currentTime  string
	deviceId     string
	deviceType   string
}

type ProviderParameters struct {
	clientId     string
	clientSecret string
	accessToken  string
	currentTime  string
	deviceId     string
	deviceType   string
}
