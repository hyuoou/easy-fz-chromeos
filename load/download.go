package load

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

func Download(url string, fileName string, downloadPath string, hash string, checkSum bool) {
	if strings.Contains(downloadPath, "$HOME") {
		strings.Replace(downloadPath, "$HOME", os.Getenv("HOME"), 1)
	}
	go func() {
		fmt.Println("Start downloading ChromeOS for the selected model name. Please wait")
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		path, err := os.Create(downloadPath + "/" + fileName + ".zip")
		if err != nil {
			log.Fatalln(err)
		}
		defer path.Close()
		io.Copy(path, resp.Body)

		if checkSum {
			fmt.Println("Check the integrity of the downloaded files")
			file, err := os.Open(downloadPath + "/" + fileName + ".zip")
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			sha1sum := sha1.New()
			if _, err := io.Copy(sha1sum, file); err != nil {
				log.Fatalln(err)
			}
			if hash == fmt.Sprintf("%x", sha1sum.Sum(nil)) {
				fmt.Println("Successful consistency check")
			} else {
				fmt.Println("Warning. Integrity check failed")
			}
		}
		os.Exit(0)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	filePath := downloadPath + "/" + fileName + ".zip"
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}
