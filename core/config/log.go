package config

type Log struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	JSONFormat bool   `mapstructure:"json_format"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackup  int    `mapstructure:"max_backup"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
