package homekit

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	ManufacturerTuya = "Tuya"
)

var (
	ErrInvalidConfigFileName = errors.New("invalid config file name")
)

type Config struct {
	Bridge  BridgeConfig   `yaml:"bridge"`
	Devices []DeviceConfig `yaml:"devices"`
}

type BridgeConfig struct {
	Name              string `yaml:"name"`
	Manufacturer      string `yaml:"manufacturer"`
	Pin               string `yaml:"pin"`
	DeviceStoragePath string `yaml:"device_storage_path"`
}

type DeviceConfig struct {
	Name         string `yaml:"name"`
	Manufacturer string `yaml:"manufacturer"`
	Id           string `yaml:"id"`
	Type         string `yaml:"type"`
}

func loadConfiguration(configFileName string) ([]byte, error) {
	if len(configFileName) < 1 {
		return nil, ErrInvalidConfigFileName
	}

	configData, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, errors.Wrap(err, "could not read config file")
	}

	return configData, nil
}

func parseConfiguration(configData []byte) (*Config, error) {
	homekitConfig := Config{}
	err := yaml.Unmarshal(configData, &homekitConfig)

	return &homekitConfig, err
}
