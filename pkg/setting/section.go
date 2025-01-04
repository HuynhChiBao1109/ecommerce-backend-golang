package setting

type Config struct {
	MySQL  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"database"`
	MaxIdleConns int    `mapstructure:"max_idle_cons"`
	MaxOpenCons  int    `mapstructure:"max_open_cons"`
	ConnMaxTime  int    `mapstructure:"con_max_life_time"`
}

type LoggerSetting struct {
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
	LogLevel   string `mapstructure:"log_level"`
	MaxBackups int    `mapstructure:"max_backups"`
}
