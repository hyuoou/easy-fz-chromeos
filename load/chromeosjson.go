package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadChromeosJson() []ChromeOS {
	fmt.Println("Loading recovery.json")
	url := "https://dl.google.com/dl/edgedl/chromeos/recovery/recovery.json"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	ChromeosJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var ChromeosDeviceList []ChromeOS
	if err := json.Unmarshal(ChromeosJson, &ChromeosDeviceList); err != nil {
		log.Fatalln(err)
	}

	ChromeosDeviceList = LoadCloudreadyJson(ChromeosDeviceList)

	return ChromeosDeviceList
}
