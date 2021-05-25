package collector

import (
	"log"
	"net"
	"os"
	"os/user"
)

type OSExporter interface {
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
