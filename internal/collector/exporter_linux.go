// +build linux

package collector

type Exporter struct{}

func (e *Exporter) GetSerialNumber() string {
	cmd := ShellCommand{
		Command:   "sudo",
		Arguments: "dmidecode -s system-serial-number",
	}
	return cmd.Execute()
}

func (e *Exporter) GetManufacturer() string {
	cmd := ShellCommand{
		Command:   "sudo",
		Arguments: "dmidecode -s system-manufacturer",
	}
	return cmd.Execute()

}

func (e *Exporter) GetSystemVersion() string {
	cmd := ShellCommand{
		Command:   "sudo",
		Arguments: "dmidecode -s system-version",
	}
	return cmd.Execute()
}
