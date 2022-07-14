package config

type Server struct {
	Port int
}

type Config struct {
	Server *Server
}

func GetConfig() *Config {
	return &Config{
		Server: &Server{
			Port: 8080,
		},
	}
}
