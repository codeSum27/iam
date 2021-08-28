package common

var Cnf *Config

type Config struct {
	debug bool `json:"debug"`
	server struct{
		Port string `json: port`
	}
	Database struct{
		Mariadb struct{
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"mariaDB"`
		Redis struct{
			Host string `json:"host"`
			Port string `json:"port"`
			Password string `json:"password"`
		} `json:"redis"`
	}
	Token struct{
		AccessSecret string `json:"ACCESS_SECRET"`
		RefreshSecret string `json:"REFRESH_SECRET"`
	} `json:"token"`
}

