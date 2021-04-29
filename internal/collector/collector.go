package collector

import (
	"inventory_old/collector"
	"log"
	"net"
	"os"
	"os/user"
)

type HostInfo struct {
	WH            string   `json:"wh"`
	UserName      string   `json:"username"`
	HostName      string   `json:"hostname"`
	MacAddress    []string `json:"macs"`
	SerialNumber  string   `json:"serial"`
	Manufacturer  string   `json:"manufacturer"`
	SystemVersion string   `json:"system_version"`
}

func BuildHostInfo(exporter Exporter, wh string) HostInfo {
	return HostInfo{
		wh,
		collector.CurrentUser(),
		collector.Hostname(),
		collector.MacAddr(),
		exporter.GetSerialNumber(),
		exporter.GetManufacturer(),
		exporter.GetSystemVersion(),
	}
}

type Exporter interface {
	GetSerialNumber() string
	GetManufacturer() string
	GetSystemVersion() string
}

func CurrentUser() string {
	username, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return username.Name
}

func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hostname
}

func MacAddr() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}
	var macs []string
	for _, ifName := range interfaces {
		mac := ifName.HardwareAddr.String()
		if mac == "" {
			continue
		}

		if !isUniq(mac, macs) {
			continue
		}

		macs = append(macs, mac)
	}
	return macs
}
