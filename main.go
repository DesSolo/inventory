package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"inventory/collector"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type PcInfo struct {
	WH string `json:"wh"`
	UserName string `json:"username"`
	HostName string `json:"hostname"`
	MacAddress []string `json:"macs"`
	SerialNumber string `json:"serial"`
	Manufacturer string `json:"manufacturer"`
	SystemVersion string `json:"system_version"`
}

func readWh()  string {
	reader := bufio.NewReader(os.Stdin)
	print("Цифры на наклейке с обратной стороны ноутбука: ")
	text, _ := reader.ReadString('\n')
	wh := strings.TrimSpace(text)
	matched, err := regexp.MatchString("WH[H\\d]\\d{6,7}", strings.ToUpper(wh))
	if err != nil {
		log.Fatalln(err)
	}
	if !matched {
		println("Неверный формат\n пример wh1234567 whh123456")
		return readWh()
	}
	return wh
}

func GetCurrentExporter() (collector.Exporter, error) {
	switch runtime.GOOS {
	case "linux":
		return &collector.Linux{}, nil
	case "windows":
		return &collector.Windows{}, nil
	default:
		return nil, errors.New("os decode error")
	}
}

func main()  {
	var silentMode bool
	flag.BoolVar(&silentMode, "s", false, "silent mode")
	flag.Parse()

	exporter, err := GetCurrentExporter()
	if err != nil {
		log.Fatalln(err)
	}
	InventoryNumber := ""
	if !silentMode {
		InventoryNumber = readWh()
	}
	data := PcInfo{
		InventoryNumber,
		collector.CurrentUser(),
		collector.Hostname(),
		collector.MacAddr(),
		exporter.GetSerialNumber(),
		exporter.GetManufacturer(),
		exporter.GetSystemVersion(),
	}
	js, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	println(string(js))
}

