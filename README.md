# HomeKit Go

Created specifically for running non HomeKit enabled devices on embedded hardware such as Ubiquiti Routers. Will compile
to the MIPS architecture.

## Setup

1. Make sure to set up provider credentials. This is currently just Tuya. Add `TUYA_CLIENT_ID` and `TUYA_CLIENT_SECRET`
   environment variables

## Tuya

Tuya makes cheap IoT devices that are rebranded by various companies. In Australia one of these rebrands is Mirabella
Genio, they make relatively very cheap and quality smart lights that
are [for sale at Kmart](https://www.kmart.com.au/product/mirabella-genio-b22-ww-wi-fi-dim-bulb/2326808) for around $13
AUD but being a Tuya device they only support Google and Amazon Home systems.

### Initial Setup

1. Download the Tuya App and Sign Up for an account
1. Add all your devices to the Tuya App

### Creating a Tuya IoT Account

1. Create a new account on iot.tuya.com and make sure you are logged in. Select United States as your country when
   signing up. This seems to skip a required verify step.
1. Go to Cloud -> Projects in the left nav drawer. If you haven't already, you will need to "purchase" the Trial Plan
   before you can proceed with this step. You will not have to add any form of payment, and the purchase is of no
   charge. Once in the Projects tab, click "Create". Make sure you select "Smart Home" for the "Industry" field and PaaS
   for the development method. Select your country of use in the for the location access option, and feel free to skip
   the services option in the next window. After you've created a new project, click into it. This will give you the
   access ID and access key we need.
1. Go to Cloud -> Project and click the project you created earlier. Then click "Link Device". Click the "Link Devices
   by App Account" tab.
1. Click "Add App Account" and scan the QR code from your smartphone / tablet app by going to the 'Me' tab in the app,
   and tapping a QR code / Scan button in the upper right. Your account will now be linked.

## Performance Tests

Coming soon - but initial testing sees this require very little resources so should be alright to run permanently on
embedded hardware

## Sources:

Resources used:

- [HomeKit Library Examples](https://github.com/brutella/hklight)
- [HomeKit Library](https://github.com/brutella/hc)
- [Docs and Setup](https://github.com/codetheweb/tuyapi/blob/master/docs/SETUP.md)
- [Tuya API - Signature](https://developer.tuya.com/en/docs/iot/singnature?id=Ka43a5mtx1gsc)
- [Tuya API - Postman](https://developer.tuya.com/en/docs/iot/set-up-postman-environment?id=Ka7o385w1svns)
- [Tuya API - Cloud IoT Explorer](https://iot.tuya.com/cloud/explorer)
