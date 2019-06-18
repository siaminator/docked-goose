package config

type Configurator interface {
	GetDriver() (string, error)
	GetDbString() (string, error)
}

type tmp struct {
	username string
	password string
	database string
	sslmode  string
	host     string
}
