package collector

import "strings"

func wmicInfo(arg string) string {
	out := RunShell(ShellCommand{"wmic", arg})
	lines := strings.Split(out, "\n")
	return lines[len(lines)-1]

}

type Windows struct {}

func (w *Windows) GetSerialNumber() string {
	return wmicInfo("bios get serialnumber")
}
func (w *Windows) GetManufacturer() string {
	return wmicInfo("computersystem get manufacturer")

}
func (w *Windows) GetSystemVersion() string {
	return wmicInfo("computersystem get model")
}