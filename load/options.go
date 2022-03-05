package load

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

func Option() Options {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		} else {
			log.Fatalln(err)
		}
	}
	return opts
}
