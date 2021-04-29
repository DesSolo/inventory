package collector

type Linux struct{}

func (l *Linux) GetSerialNumber() string {
	cmd := ShellCommand{"sudo", "dmidecode -s system-serial-number"}
	return RunShell(cmd)
}

func (l *Linux) GetManufacturer() string {
	cmd := ShellCommand{"sudo", "dmidecode -s system-manufacturer"}
	return RunShell(cmd)

}

func (l *Linux) GetSystemVersion() string {
	cmd := ShellCommand{"sudo", "dmidecode -s system-version"}
	return RunShell(cmd)
}
