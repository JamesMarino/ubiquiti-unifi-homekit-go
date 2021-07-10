package homekit

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/jamesmarino/ubiquiti-unifi-homekit-go/pkg/providers/tuya"
)

func InitialiseConfiguration(configFilename string) (*Config, error) {
	config, err := loadConfiguration(configFilename)
	if err != nil {
		return nil, err
	}

	return parseConfiguration(config)
}

func BroadcastDevices(config *Config) error {
	var homeAccessories []*accessory.Accessory

	for _, device := range config.Devices {
		accessoryInfo := accessory.Info{
			Name:         device.Name,
			Manufacturer: device.Manufacturer,
		}

		homekitAccessory := accessory.NewLightbulb(accessoryInfo)

		switch device.Manufacturer {
		case ManufacturerTuya:
			tuyaRequest := tuya.NewTuya(device.Id, device.Type)
			homekitAccessory.Lightbulb.On.OnValueRemoteUpdate(tuyaRequest.ToggleDevice)
		}

		homeAccessories = append(homeAccessories, homekitAccessory.Accessory)
	}

	bridgeInfo := accessory.Info{
		Name:         config.Bridge.Name,
		Manufacturer: config.Bridge.Manufacturer,
	}

	bridgeAccessory := accessory.NewBridge(bridgeInfo)

	deviceConfig := hc.Config{
		StoragePath: config.Bridge.DeviceStoragePath,
		Pin:         config.Bridge.Pin,
	}

	ipTransport, err := hc.NewIPTransport(deviceConfig, bridgeAccessory.Accessory, homeAccessories...)
	if err != nil {
		return err
	}

	hc.OnTermination(func() {
		<-ipTransport.Stop()
	})

	ipTransport.Start()

	return nil
}
