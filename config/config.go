package config

type Config struct {
	LogLevel             int `json:"loglevel"`
	AppPort              int `json:"port"`
	HttpRequestTimeout   int `json:"httpRequestTimeout"`
	HttpConnectinTimeout int `json:"httpConnectionTimeout"`
}
