package load

const (
	AppName    = "easy-fz-chromeos"
	AppVersion = "0.1.0"
)

type Options struct {
	Version  bool `short:"v" long:"version" description:"Show version"`
	Download bool `short:"d" long:"download" description:"Open a browser and download the file"`
}

type ChromeOS struct {
	Channel       string `json:"channel"`
	File          string `json:"file"`
	Filesize      int    `json:"filesize"`
	Hwidmatch     string `json:"hwidmatch"`
	Manufacturer  string `json:"manufacturer"`
	Md5           string `json:"md5"`
	Model         string `json:"model"`
	Name          string `json:"name"`
	Sha1          string `json:"sha1"`
	Url           string `json:"url"`
	Version       string `json:"version"`
	Zipfilesize   int    `json:"zipfilesize"`
	ChromeVersion string `json:"chrome_version"`
}
