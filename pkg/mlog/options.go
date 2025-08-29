package mlog

type LoggerOptions struct {
	Level      string `yaml:"level"`
	LogsDir    string `yaml:"dir"`
	MaxSize    int    `yaml:"maxsize"`
	MaxAge     int    `yaml:"maxage"`
	MaxBackups int    `yaml:"maxbackups"`
	Compress   bool   `yaml:"compress"`
	Console    bool   `yaml:"console"`
}
