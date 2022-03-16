package load

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

func Option(downloadCheck *bool, downloadPath *string) Options {
	var opts Options
	opts.Download = func(s string) {
		// 引数で与えられたディレクトリが存在するか確認
		_, err := os.Stat(s)
		if os.IsNotExist(err) {
			log.Fatalln("Directory does not exist.")
		}

		*downloadCheck = true
		*downloadPath = s
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = AppName
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
