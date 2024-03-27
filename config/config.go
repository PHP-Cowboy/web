package config

type ServerConfig struct {
	Port      int         `mapstructure:"port" json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	JwtInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	Sms       Sms         `mapstructure:"sms" json:"sms"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type Sms struct {
	AppKey string `mapstructure:"appKey" json:"appKey"`
	Secret string `mapstructure:"secret" json:"secret"`
}
