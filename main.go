package main

import (
	"fmt"

	"github.com/hyuoou/easy-fz-chromeos/load"
)

func main() {
	chromeosList := load.LoadJson()

	for i := range chromeosList {
		fmt.Printf("%s: %s\n", chromeosList[i].Model, chromeosList[i].Url)
	}
}
