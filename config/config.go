package config

type Settings struct {
	S3BasedURL string
	BaseURL    string
}

var ConfigSettings Settings

func SetConfig(s3basedUrl, baseUrl string) {
	ConfigSettings.S3BasedURL = s3basedUrl
	ConfigSettings.BaseURL = baseUrl
}
