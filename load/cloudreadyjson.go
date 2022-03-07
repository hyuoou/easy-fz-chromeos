package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadCloudreadyJson(ChromeosDeviceList []ChromeOS) []ChromeOS {
	fmt.Println("Loadint cloudready_recovery.json")
	url := "https://dl.google.com/dl/edgedl/chromeos/recovery/cloudready_recovery.json"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	CloudreadyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var CloudreadyList []ChromeOS
	if err := json.Unmarshal(CloudreadyJson, &CloudreadyList); err != nil {
		log.Fatalln(err)
	}

	ChromeosDeviceList = append(ChromeosDeviceList, CloudreadyList...)

	return ChromeosDeviceList
}
