package store

type Config struct {
	DatabaseHostName string
}

func NewConfig(dbhostname string) *Config {
	return &Config{
		DatabaseHostName: dbhostname,
	}
}
