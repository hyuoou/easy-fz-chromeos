package load

const (
	AppVersion = "0.0.2"
	AppName    = "easy-fz-chromeos"
)

type Options struct {
	Version  bool         `short:"v" long:"version" description:"Show version"`
	Download func(string) `short:"d" long:"download" description:"Download ChromeOS to the specified path"`
}
