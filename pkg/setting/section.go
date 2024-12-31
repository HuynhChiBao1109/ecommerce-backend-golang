package setting

type Config struct {
	MySQL MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	MaxIdle     int    `mapstructure:"maxidle"`
	MaxOpen     int    `mapstructure:"maxopen"`
	ConnMaxTime int    `mapstructure:"connmaxtime"`
}
