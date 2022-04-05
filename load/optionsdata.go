package load

const (
	AppVersion = "0.0.2"
	AppName    = "easy-fz-chromeos"
)

type Options struct {
	Version  bool         `short:"v" long:"version" description:"Show version"`
	CheckSum bool         `short:"c" long:"checksum" description:"Check the integrity of downloaded files"`
	Download func(string) `short:"d" long:"download" description:"Download ChromeOS to the specified path"`
}
