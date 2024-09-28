package config

type ServerConfig struct {
	Port      int         `json:"port"`
	MysqlInfo MysqlConfig `json:"mysqlInfo"`
	JwtInfo   JWTConfig   `json:"jwtInfo"`
	Sms       Sms         `json:"sms"`
	ThirdApp  ThirdApp    `json:"thirdApp"`
	AliCloud  AliCloud    `json:"aliCloud"`
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
	ExpiresHour int    `json:"expiresHour"`
	AddHour     int    `json:"addHour"`
}

type Sms struct {
	AppKey string `json:"appKey"`
	Secret string `json:"secret"`
}

type AliCloud struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
}

type ThirdApp struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}
