package load

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func Download(url string, fileName string, downloadPath string) {
	go func() {
		fmt.Println("Start downloading ChromeOS for the selected model name. Please wait")
		err := exec.Command("wget", "-P", downloadPath, url).Run()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done")
		os.Exit(0)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	filePath := downloadPath + "/" + fileName + ".zip"
	exec.Command("rm", filePath).Run()
	os.Exit(1)
}
