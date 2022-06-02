package load

import (
	"fmt"

	"github.com/ktr0731/go-fuzzyfinder"
)

func Finder(chromeosList []ChromeOS) (int, error) {
	idx, err := fuzzyfinder.Find(
		chromeosList,
		func(i int) string {
			return chromeosList[i].Model
		},
		fuzzyfinder.WithPreviewWindow(func(i, width, height int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Channel       : %s\nFile          : %s\nFilesize      : %vB\nHwidmatch     : %s\nManufacturer  : %s\nMd5           : %s\nModel         : %s\nName          : %s\nSha1          : %s\nVersion       : %s\nZipfilesize   : %vB\nchromeosList  : %s\n",
				chromeosList[i].Channel,
				chromeosList[i].File,
				chromeosList[i].Filesize,
				chromeosList[i].Hwidmatch,
				chromeosList[i].Manufacturer,
				chromeosList[i].Md5,
				chromeosList[i].Model,
				chromeosList[i].Name,
				chromeosList[i].Sha1,
				chromeosList[i].Version,
				chromeosList[i].Zipfilesize,
				chromeosList[i].ChromeVersion)
		}),
	)
	return idx, err
}
