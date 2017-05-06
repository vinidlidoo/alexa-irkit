package main

import "fmt"

// spec.go defines Alexa Smart Home Skill request and response.
// Currently, it only defines what **need** (so apparently, API
// is not completed).
//
// See more at Alexa skill Kit documentation.
// https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/smart-home-skill-api-reference

type NameSpace string

const (
	AlexaConnectedHomeDiscovery NameSpace = "Alexa.ConnectedHome.Discovery"
	AlexaConnectedHomeControl   NameSpace = "Alexa.ConnectedHome.Control"
)

type Name string

const (
	DiscoverAppliancesRequest  Name = "DiscoverAppliancesRequest"
	DiscoverAppliancesResponse Name = "DiscoverAppliancesResponse"

	TurnOnRequest      Name = "TurnOnRequest"
	TurnOnConfirmation Name = "TurnOnConfirmation"

	TurnOffRequest      Name = "TurnOffRequest"
	TurnOffConfirmation Name = "TurnOffConfirmation"

	IncrementPercentageRequest      Name = "IncrementPercentageRequest"
	IncrementPercentageConfirmation Name = "IncrementPercentageConfirmation"

	DecrementPercentageRequest      Name = "DecrementPercentageRequest"
	DecrementPercentageConfirmation Name = "DecrementPercentageConfirmation"
)

type ApplianceId string

const (
	irkitLights   ApplianceId = "irkitLights"
	irkitAC       ApplianceId = "irkitAC"
	irkitRoomba   ApplianceId = "irkitRoomba"
	irkitDimLight ApplianceId = "irkitDimLight"
)

const (
	DriverInternalError     = "DriverInternalError"
	InvalidAccessTokenError = "InvalidAccessTokenError"
)

const (
	PayloadVersion = "2"
)

type Directive struct {
	Header  Header                 `json:"header"`
	Payload map[string]interface{} `json:"payload"`
	//Payload Payload `json:"payload"`
}

// Header has a set of expected fields that are the same across message types.
type Header struct {
	MessageID string    `json:"messageId"`
	Name      Name      `json:"name"`
	NameSpace NameSpace `json:"namespace"`
	Version   string    `json:"payloadVersion"`
}

//type Payload struct {
//	AccessToken          string                `json:"accessToken"`
//	DiscoveredAppliances []DiscoveredAppliance `json:"appliance"`
//}

type DiscoveredAppliance struct {
	ApplianceId         ApplianceId `json:"applianceId"`
	ManufacturerName    string      `json:"manufacturerName"`
	ModelName           string      `json:"modelName"`
	Version             string      `json:"version"`
	FriendlyName        string      `json:"friendlyName"`
	FriendlyDescription string      `json:"friendlyDescription"`
	IsReachable         bool        `json:"isReachable"`
	Actions             []string    `json:"actions"`

	AdditionalApplianceDetails map[string]string `json:"additionalApplianceDetails"`
}

func (d *Directive) AccessToken() (string, error) {
	payload := d.Payload
	v, ok := payload["accessToken"]
	if !ok {
		return "", fmt.Errorf("missing access token")
	}

	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("failed to type assertion")
	}

	return token, nil
}

func (d *Directive) AccessApplianceId() (string, error) {
	payload := d.Payload
	appliance, ok := payload["appliance"]
	if !ok {
		return "", fmt.Errorf("missing appliance")
	}
	applianceMap, ok := appliance.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("failed to type assertion")
	}
	applianceId, ok := applianceMap["applianceId"]
	if !ok {
		return "", fmt.Errorf("missing appliance Id")
	}
	id, ok := applianceId.(string)
	if !ok {
		return "", fmt.Errorf("failed to type assertion")
	}
	return id, nil
}

func (d *Directive) AccessDeltaPercentage() (float64, error) {
	payload := d.Payload
	deltaPercentage, ok := payload["deltaPercentage"]
	if !ok {
		return 0.0, fmt.Errorf("missing delta percentage")
	}
	deltaMap, ok := deltaPercentage.(map[string]interface{})
	if !ok {
		return 0.0, fmt.Errorf("failed to type assertion")
	}
	v, ok := deltaMap["value"]
	if !ok {
		return 0.0, fmt.Errorf("missing delta percentage value")
	}
	delta, ok := v.(float64)
	if !ok {
		return 0.0, fmt.Errorf("failed to type assertion")
	}
	return delta, nil
}

func ErrorResponse(name Name, messageId string, namespace NameSpace) *Directive {
	return &Directive{
		Header: Header{
			MessageID: messageId,
			Name:      name,
			NameSpace: namespace,
			Version:   PayloadVersion,
		},

		Payload: map[string]interface{}{},
		//Payload: Payload{},
	}
}
