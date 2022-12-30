package config

type ServerConfig struct {
	Port      int         `mapstructure:"port" json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	JwtInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	RocketMQ  string      `mapstructure:"rocket_mq" json:"rocket_mq"`
	Odbc      string      `mapstructure:"odbc" json:"odbc"`
	ESign     ESign       `mapstructure:"e_sign" json:"e_sign"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password"`
	Expire   int    `mapstructure:"expire" json:"expire"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ESign struct {
	ESignHost      string `json:"e_sign_host"`
	ESignProjectId string `json:"e_sign_project_id"`
	ESignSecret    string `json:"e_sign_secret"`
	ESignCallback  string `json:"e_sign_callback"`
}
