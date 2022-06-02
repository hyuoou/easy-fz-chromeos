package load

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Download(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Run()
	case "windows":
		err = exec.Command("start", url).Run()
	case "darwin":
		err = exec.Command("open", url).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatalln(err)
	}
}
