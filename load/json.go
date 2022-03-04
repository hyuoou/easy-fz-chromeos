package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func LoadChromeosJson() []ChromeOS {
	fmt.Println("recovery.jsonをダウンロードしています。")
	err := exec.Command("wget", "-Orecovery.json", "https://dl.google.com/dl/edgedl/chromeos/recovery/recovery.json").Run()
	if err != nil {
		log.Fatalln(err)
	}

	ChromeosJson, err := ioutil.ReadFile("recovery.json")
	if err != nil {
		log.Fatalln(err)
	}
	var ChromeosDeviceList []ChromeOS
	if err := json.Unmarshal(ChromeosJson, &ChromeosDeviceList); err != nil {
		log.Fatalln(err)
	}

	err = exec.Command("rm", "recovery.json").Run()
	if err != nil {
		log.Fatalln(err)
	}

	return ChromeosDeviceList
}
