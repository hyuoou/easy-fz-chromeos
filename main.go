package main

import (
	"fmt"
	"log"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	ChromeosDeviceList := load.LoadChromeosJson()
	ChromeosDeviceList = load.LoadCloudreadyJson(ChromeosDeviceList)

	idx, err := load.ShowFuzzyfinder(ChromeosDeviceList)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s %s\n", ChromeosDeviceList[idx].Model, ChromeosDeviceList[idx].Url)
}
