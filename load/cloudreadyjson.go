package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func LoadCloudreadyJson(ChromeosDeviceList []ChromeOS) []ChromeOS {
	fmt.Println("Downloadng cloudready_recovery.json")
	err := exec.Command("wget", "-Ocloudready_recovery.json", "https://dl.google.com/dl/edgedl/chromeos/recovery/cloudready_recovery.json").Run()
	if err != nil {
		log.Fatalln(err)
	}

	CloudreadyJson, err := ioutil.ReadFile("cloudready_recovery.json")
	if err != nil {
		log.Fatalln(err)
	}
	var CloudreadyList []ChromeOS
	if err := json.Unmarshal(CloudreadyJson, &CloudreadyList); err != nil {
		log.Fatalln(err)
	}

	ChromeosDeviceList = append(ChromeosDeviceList, CloudreadyList...)

	err = exec.Command("rm", "cloudready_recovery.json").Run()
	if err != nil {
		log.Fatalln(err)
	}

	return ChromeosDeviceList
}
