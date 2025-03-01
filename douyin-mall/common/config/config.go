package config

type Config struct {
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Server struct {
		Port int
	}
}

var GlobalConfig Config

func Init() error {
	// TODO: 实现配置加载逻辑
	return nil
}
