// +build windows

package collector

import "strings"

func wmicInfo(arg string) string {
	cmd := ShellCommand{
		Command:   "wmic",
		Arguments: arg,
	}
	out := cmd.Execute()
	lines := strings.Split(out, "\n")
	return lines[len(lines)-1]
}

type Exporter struct{}

func (e *Exporter) GetSerialNumber() string {
	return wmicInfo("bios get serialnumber")
}

func (e *Exporter) GetManufacturer() string {
	return wmicInfo("computersystem get manufacturer")

}

func (e *Exporter) GetSystemVersion() string {
	return wmicInfo("computersystem get model")
}
