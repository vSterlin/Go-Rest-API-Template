package config

type Config struct {
	Server
	Store
}

type Server struct {
	BindAddr string
}

type Store struct {
	Driver string
	DSN    string
}

func New() (*Config, error) {
	// TODO
	return &Config{}, nil
}

func Defaults() *Config {
	return &Config{
		Server: Server{
			BindAddr: ":8080",
		},
		Store: Store{
			Driver: "postgres",
			DSN:    "user=v dbname=template sslmode=disable",
		},
	}
}
