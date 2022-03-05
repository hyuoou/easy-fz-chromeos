package load

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func Download(url string, fileName string) {
	go func() {
		fmt.Println("Start downloading ChromeOS for the selected model name. Please wait")
		err := exec.Command("wget", url).Run()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done")
		os.Exit(0)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	exec.Command("rm", fileName+".zip").Run()
	os.Exit(1)
}
