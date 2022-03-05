package main

import (
	"fmt"
	"log"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	opts := load.Option()

	if opts.Version {
		fmt.Println(opts.Version)
		return
	}

	ChromeosDeviceList := load.LoadChromeosJson()

	idx, err := load.ShowFuzzyfinder(ChromeosDeviceList)
	if err != nil {
		log.Fatalln(err)
	}

	if opts.Download {
		load.Download(ChromeosDeviceList[idx].Url, ChromeosDeviceList[idx].File)
	}

	fmt.Printf("%s %s\n", ChromeosDeviceList[idx].Model, ChromeosDeviceList[idx].Url)
}
