package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/hyuoou/easy-fz-chromeos/load"
	"github.com/jessevdk/go-flags"
)

const (
	version = "0.0.1"
	appName = "easy-fz-chromeos"
)

type Options struct {
	Version  bool `short:"v" long:"version" description:"Show version"`
	Download bool `short:"d" long:"download" description:"Download ChromeOS for the selected model name"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		if flags.WroteHelp(err) {
			return
		} else {
			log.Fatalln(err)
		}
	}

	if opts.Version {
		fmt.Println(version)
		return
	}

	ChromeosDeviceList := load.LoadChromeosJson()

	idx, err := load.ShowFuzzyfinder(ChromeosDeviceList)
	if err != nil {
		log.Fatalln(err)
	}

	if opts.Download {
		go func() {
			fmt.Println("Start downloading ChromeOS for the selected model name. Please wait")
			err := exec.Command("wget", ChromeosDeviceList[idx].Url).Run()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Done")
			return
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		fileName := ChromeosDeviceList[idx].File + ".zip"
		exec.Command("rm", fileName).Run()
		os.Exit(1)
	}

	fmt.Printf("%s %s\n", ChromeosDeviceList[idx].Model, ChromeosDeviceList[idx].Url)
}
