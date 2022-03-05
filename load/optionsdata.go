package load

const (
	version = "0.0.1"
	appName = "easy-fz-chromeos"
)

type Options struct {
	Version  bool `short:"v" long:"version" description:"Show version"`
	Download bool `short:"d" long:"download" description:"Download ChromeOS for the selected model name"`
}
