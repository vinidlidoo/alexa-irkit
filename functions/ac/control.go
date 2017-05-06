package main

import (
	"context"
	"fmt"
	irkit "github.com/tcnksm/go-irkit/v1"
	"math"
	"os"
	"time"
)

// handleControl handles TurnON/TurnOFF requests
// It sends AC on/off signal via IRKit internet HTTP API
func handleControl(d *Directive) (*Directive, error) {

	irClientKey := os.Getenv(EnvIRClientKey)
	if len(irClientKey) == 0 {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("missing %q env var", EnvIRClientKey)
	}

	irDeviceId := os.Getenv(EnvIRDeviceID)
	if len(irDeviceId) == 0 {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("missing %q env var", EnvIRDeviceID)
	}

	// Construct IRKit internet client
	client := irkit.DefaultInternetClient()
	// Access Appliance Id

	id, err := d.AccessApplianceId()
	if err != nil {
		res := ErrorResponse(
			DriverInternalError,
			d.Header.MessageID,
			d.Header.NameSpace)
		return res, err
	}

	switch applianceId := ApplianceId(id); applianceId {
	case irkitLights:
		switch name := d.Header.Name; name {
		case TurnOnRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgLightsOn)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOnConfirmation
			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case TurnOffRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgLightsOff)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOffConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case IncrementPercentageRequest:
			delta, err := d.AccessDeltaPercentage()
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}
			incr := int(math.Trunc(delta / 25.0))
			if incr > 0 {
				for i := 0; i < incr; i++ {
					err = client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgLightsIncr)
					if err != nil {
						res := ErrorResponse(DriverInternalError, "", "")
						return res, err
					}
				}
			}
			// Change header name
			d.Header.Name = IncrementPercentageConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case DecrementPercentageRequest:
			delta, err := d.AccessDeltaPercentage()
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}
			incr := int(math.Trunc(delta / 25.0))
			if incr > 0 {
				for i := 0; i < incr; i++ {
					err = client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgLightsDecr)
					if err != nil {
						res := ErrorResponse(DriverInternalError, "", "")
						return res, err
					}
				}
			}
			// Change header name
			d.Header.Name = DecrementPercentageConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		default:
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("unexpected name: %s", name)
		}
	case irkitAC:
		switch name := d.Header.Name; name {
		case TurnOnRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgACOn)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOnConfirmation
			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case TurnOffRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgACOff)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOffConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil
		default:
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("unexpected control command: %s", name)
		}
	case irkitRoomba:
		switch name := d.Header.Name; name {
		case TurnOnRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgRoombaToggle)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOnConfirmation
			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case TurnOffRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgRoombaToggle)
			time.Sleep(1500 * time.Millisecond)
			err = client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgRoombaDock)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOffConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil
		default:
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("unexpected control command: %s", name)
		}
	case irkitDimLight:
		switch name := d.Header.Name; name {
		case TurnOnRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgDimLightOn)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOnConfirmation
			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil

		case TurnOffRequest:
			err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgLightsOff)
			if err != nil {
				res := ErrorResponse(DriverInternalError, "", "")
				return res, err
			}

			// Change header name
			d.Header.Name = TurnOffConfirmation

			return &Directive{
				Header:  d.Header,
				Payload: map[string]interface{}{},
			}, nil
		default:
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("unexpected control command: %s", name)
		}
	default:
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("unexpected applianceId: %q", applianceId)
	}
}
