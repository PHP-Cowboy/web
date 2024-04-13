package config

type ServerConfig struct {
	Port      int         `json:"port"`
	MysqlInfo MysqlConfig `json:"mysqlInfo"`
	JwtInfo   JWTConfig   `json:"jwtInfo"`
	Sms       Sms         `json:"sms"`
	ThirdApp  ThirdApp    `json:"thirdApp"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type JWTConfig struct {
	SigningKey  string `json:"key"`
	ExpiresHour string `json:"expiresHour"`
}

type Sms struct {
	AppKey string `json:"appKey"`
	Secret string `json:"secret"`
}

type ThirdApp struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}
