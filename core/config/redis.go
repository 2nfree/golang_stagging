package config

type Redis struct {
	Db       int    `mapstructure:"db"`
	IP       string `mapstructure:"ip"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}
