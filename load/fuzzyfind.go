package load

import (
	"fmt"

	"github.com/ktr0731/go-fuzzyfinder"
)

func ShowFuzzyfinder(ChromeosDeviceList []ChromeOS) (int, error) {
	idx, err := fuzzyfinder.Find(
		ChromeosDeviceList,
		func(i int) string {
			return ChromeosDeviceList[i].Model
		},
		fuzzyfinder.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Channel       : %s\nFile          : %s\nFilesize      : %vB\nHwidmatch     : %s\nManufacturer  : %s\nMd5           : %s\nModel         : %s\nName          : %s\nSha1          : %s\nVersion       : %s\nZipfilesize   : %vB\nChromeVersion : %s\n",
				ChromeosDeviceList[i].Channel,
				ChromeosDeviceList[i].File,
				ChromeosDeviceList[i].Filesize,
				ChromeosDeviceList[i].Hwidmatch,
				ChromeosDeviceList[i].Manufacturer,
				ChromeosDeviceList[i].Md5,
				ChromeosDeviceList[i].Model,
				ChromeosDeviceList[i].Name,
				ChromeosDeviceList[i].Sha1,
				ChromeosDeviceList[i].Version,
				ChromeosDeviceList[i].Zipfilesize,
				ChromeosDeviceList[i].ChromeVersion)
		}),
	)
	return idx, err
}
