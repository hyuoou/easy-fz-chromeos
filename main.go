package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	opts := load.FlagParse()

	chromeosList := load.LoadJson()
	idx, err := load.Finder(chromeosList)
	if err != nil {
		log.Fatalln(err)
	}

	if opts.Download {
		load.Download(chromeosList[idx].Url)
		os.Exit(0)
	}

	fmt.Printf("%s: %s\n", chromeosList[idx].Model, chromeosList[idx].Url)
}
