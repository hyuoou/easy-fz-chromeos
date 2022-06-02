package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadJson() []ChromeOS {
	urls := []string{
		"https://dl.google.com/dl/edgedl/chromeos/recovery/recovery.json",
		"https://dl.google.com/dl/edgedl/chromeos/recovery/cloudready_recovery.json",
	}
	var chromeosList, cloudreadyList []ChromeOS

	fmt.Println("Loading json")

	for idx, url := range urls {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		chromeosJson, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if idx == 0 {
			if err := json.Unmarshal(chromeosJson, &chromeosList); err != nil {
				log.Fatalln(err)
			}
		} else {
			if err := json.Unmarshal(chromeosJson, &cloudreadyList); err != nil {
				log.Fatalln(err)
			}
			chromeosList = append(chromeosList, cloudreadyList...)
		}
	}

	return chromeosList
}
