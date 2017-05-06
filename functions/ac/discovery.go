package main

import "fmt"

// handleDiscovery handles discovery appliances request.
// Currently, it only returns constant value.
func handleDiscovery(d *Directive) (*Directive, error) {
	if name := d.Header.Name; name != DiscoverAppliancesRequest {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("unexpected name: %s", name)
	}

	// Change header name
	d.Header.Name = DiscoverAppliancesResponse

	// Construct response
	res := &Directive{
		Header: d.Header,
		Payload: map[string]interface{}{
			//Payload: Payload{
			"discoveredAppliances": []DiscoveredAppliance{
				{
					ApplianceId:         "irkitLights",
					ManufacturerName:    "vethier",
					ModelName:           "irkit-001",
					Version:             "0.0.1",
					FriendlyName:        "Lights",
					FriendlyDescription: "Lights",
					IsReachable:         true,
					Actions: []string{
						"turnOn",
						"turnOff",
						"incrementPercentage",
						"decrementPercentage",
					},
					AdditionalApplianceDetails: map[string]string{},
				},
				{
					ApplianceId:         "irkitAC",
					ManufacturerName:    "vethier",
					ModelName:           "irkit-001",
					Version:             "0.0.1",
					FriendlyName:        "AC",
					FriendlyDescription: "AirCon",
					IsReachable:         true,
					Actions: []string{
						"turnOn",
						"turnOff",
					},
					AdditionalApplianceDetails: map[string]string{},
				},
				{
					ApplianceId:         "irkitRoomba",
					ManufacturerName:    "vethier",
					ModelName:           "irkit-001",
					Version:             "0.0.1",
					FriendlyName:        "Roomba",
					FriendlyDescription: "Roomba",
					IsReachable:         true,
					Actions: []string{
						"turnOn",
						"turnOff",
					},
					AdditionalApplianceDetails: map[string]string{},
				},
				{
					ApplianceId:         "irkitDimLight",
					ManufacturerName:    "vethier",
					ModelName:           "irkit-001",
					Version:             "0.0.1",
					FriendlyName:        "Dim",
					FriendlyDescription: "Dim",
					IsReachable:         true,
					Actions: []string{
						"turnOn",
						"turnOff",
					},
					AdditionalApplianceDetails: map[string]string{},
				},
			},
		},
	}

	return res, nil
}
