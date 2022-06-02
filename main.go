package main

import (
	"fmt"
	"log"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	chromeosList := load.LoadJson()

	idx, err := load.Finder(chromeosList)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s: %s\n", chromeosList[idx].Model, chromeosList[idx].Url)
}
