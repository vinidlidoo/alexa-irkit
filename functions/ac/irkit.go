package main

import (
	"encoding/json"
	irkit "github.com/tcnksm/go-irkit/v1"
)

const (
// IRKit related environmental var
//EnvIRClientKey = "IR_CLIENT_KEY"
//EnvIRDeviceID  = "IR_DEVICE_ID"
)

var (
	IRKitMsgLightsOn     irkit.Message
	IRKitMsgLightsOff    irkit.Message
	IRKitMsgACOn         irkit.Message
	IRKitMsgACOff        irkit.Message
	IRKitMsgRoombaToggle irkit.Message
	IRKitMsgRoombaDock   irkit.Message
	IRKitMsgLightsIncr   irkit.Message
	IRKitMsgLightsDecr   irkit.Message
	IRKitMsgDimLightOn   irkit.Message
)

// readSigal reads IRKit signal json data from ./signals directory.
// Files should be transformed into go binary data by go-bindata.
func readSignal(path string, msg *irkit.Message) error {
	data, err := Asset(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, msg)
}

func init() {
	err := readSignal("signals/lights-on.json", &IRKitMsgLightsOn)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/lights-off.json", &IRKitMsgLightsOff)
	if err != nil {
		panic(err)
	}
	err = readSignal("signals/ac-on.json", &IRKitMsgACOn)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/ac-off.json", &IRKitMsgACOff)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/roomba-toggle.json", &IRKitMsgRoombaToggle)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/roomba-dock.json", &IRKitMsgRoombaDock)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/lights-incr.json", &IRKitMsgLightsIncr)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/lights-decr.json", &IRKitMsgLightsDecr)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/dimlight-on.json", &IRKitMsgDimLightOn)
	if err != nil {
		panic(err)
	}
}
