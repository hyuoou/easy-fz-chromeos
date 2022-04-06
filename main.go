package main

import (
	"fmt"
	"log"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	DownloadCheck := false
	DownloadPath := ""
	// 引数を解析
	opts := load.Option(&DownloadCheck, &DownloadPath)

	if opts.Version {
		fmt.Println(load.AppVersion)
		return
	}

	ChromeosDeviceList := load.LoadChromeosJson()

	idx, err := load.ShowFuzzyfinder(ChromeosDeviceList)
	if err != nil {
		log.Fatalln(err)
	}

	if DownloadCheck {
		load.Download(ChromeosDeviceList[idx].Url, ChromeosDeviceList[idx].File, DownloadPath, ChromeosDeviceList[idx].Sha1)
	}

	fmt.Printf("%s %s\n", ChromeosDeviceList[idx].Model, ChromeosDeviceList[idx].Url)
}
