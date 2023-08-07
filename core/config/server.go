package config

type Server struct {
	Listen  string `mapstructure:"listen"`
	RunMode string `mapstructure:"run_mode"`
}
