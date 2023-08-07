package config

type SMTP struct {
	StartTLS     bool   `mapstructure:"start_tls"`
	MaxClient    int    `mapstructure:"max_client"`
	SMTPHost     string `mapstructure:"smtp_host"`
	SMTPPort     string `mapstructure:"smtp_port"`
	SMTPUsername string `mapstructure:"smtp_username"`
	SMTPPassword string `mapstructure:"smtp_password"`
}
