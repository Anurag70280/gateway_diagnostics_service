package models

type Device struct {
	SerialNumber   string  `json:"serialNumber"`
	Name           string  `json:"name"`
	SiteName       string  `json:"siteName"`
	SiteId         int     `json:"siteId"`
	DeviceType     string  `json:"deviceType"`
	FwVer          string  `json:"fwVer"`
	Status         string  `json:"status"`
	BatteryVoltage float32 `json:"batteryVoltage"`
	AccessPointId  int     `json:"accessPointId"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}

type GetDevicesResponse struct {
	Type    string  `json:"type"`
	Message Devices `json:"message"`
}
