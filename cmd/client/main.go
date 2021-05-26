package main

import (
	"bufio"
	"flag"
	"inventory/internal/collector"
	"inventory/internal/uploader"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	serverURL = "http://127.0.0.1:8090/api/v1/client"
	token     = "inventory"
)

func readWh() string {
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

func GetCurrentExporter() (collector.OSExporter, error) {
	return &collector.Exporter{}, nil
}

func main() {
	var silentMode bool
	flag.BoolVar(&silentMode, "s", false, "silent mode")
	flag.Parse()

	exporter, err := GetCurrentExporter()
	if err != nil {
		log.Fatalln(err)
	}

	var InventoryNumber string
	if !silentMode {
		InventoryNumber = readWh()
	}

	data := collector.CollectHostInfo(exporter, InventoryNumber)
	up := uploader.NewRest(serverURL, token)
	if err := up.Upload(&data); err != nil {
		log.Fatalln(err)
	}
}
