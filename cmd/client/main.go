package main

import (
	"bufio"
	"errors"
	"flag"
	"inventory/internal/collector"
	"inventory/internal/storage"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

const (
	serverURL = "http://127.0.0.1"
	username  = "inventory"
	password  = "Kmx4r9d0c@!"
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

	data := collector.BuildHostInfo(exporter, InventoryNumber)
	st := storage.NewRest(serverURL, username, password)
	if err := st.Send(&data); err != nil {
		log.Fatalln(err)
	}
}
