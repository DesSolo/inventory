package collector

import "errors"

type HostInfo struct {
	WH            string   `json:"wh"`
	UserName      string   `json:"username"`
	HostName      string   `json:"hostname"`
	MacAddress    []string `json:"macs"`
	SerialNumber  string   `json:"serial"`
	Manufacturer  string   `json:"manufacturer"`
	SystemVersion string   `json:"system_version"`
}

func (h *HostInfo) Validate() error {
	if h.UserName == "" {
		return errors.New("value username is empty")
	}
	if h.HostName == "" {
		return errors.New("value hostname is empty")
	}
	if len(h.MacAddress) == 0 {
		return errors.New("value macs is empty")
	}
	if h.SerialNumber == "" {
		return errors.New("value serial is empty")
	}
	if h.Manufacturer == "" {
		return errors.New("value manufacturer is empty")
	}
	if h.SystemVersion == "" {
		return errors.New("value system_version is empty")
	}
	return nil
}

func CollectHostInfo(ex OSExporter, wh string) HostInfo {
	return HostInfo{
		WH:            wh,
		UserName:      CurrentUser(),
		HostName:      Hostname(),
		MacAddress:    MacAddr(),
		SerialNumber:  ex.GetSerialNumber(),
		Manufacturer:  ex.GetManufacturer(),
		SystemVersion: ex.GetSystemVersion(),
	}
}
