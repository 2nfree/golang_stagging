package config

type Database struct {
	LogMode       string `mapstructure:"log_mode"`
	ConsoleLog    bool   `mapstructure:"console_log"`
	SlowThreshold int    `mapstructure:"slow_threshold"`
	Path          string `mapstructure:"path"`
	Mysql         Mysql  `mapstructure:"mysql"`
}

type Mysql struct {
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Hostname    string `mapstructure:"hostname"`
	Database    string `mapstructure:"database"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxLifeTime int    `mapstructure:"max_life_time"`
}
