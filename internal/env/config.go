package env

import "rcbs/internal/values"

type config struct {
	Mode values.Mode `env:"MODE" envDefault:"unset"`

	// Database
	Mongo struct {
		Host       string `env:"HOST" envDefault:"localhost"`
		Port       string `env:"PORT" envDefault:"27017"`
		User       string `env:"USER" envDefault:""`
		Pass       string `env:"PASS" envDefault:""`
		Database   string `env:"DATABASE" envDefault:"rcbs"`
		Additional string `env:"ADDITIONAL" envDefault:""`
	} `envPrefix:"MONGO_"`

	// Server
	Server struct {
		Port    int    `env:"PORT" envDefault:"8080"`
		BaseURI string `env:"BASE_URI" envDefault:"http://localhost:8080"`

		// Mail
		Mail struct {
			RelayHost          string `env:"RELAY_HOST"`
			RelayPort          int    `env:"RELAY_PORT"`
			RelayUsername      string `env:"RELAY_USERNAME"`
			RelayPassword      string `env:"RELAY_PASSWORD"`
			RelayTLS           bool   `env:"RELAY_TLS"`
			RelayIgnoreTLSCert bool   `env:"IGNORE_TLS_CERT"`
			From               string `env:"FROM" envDefault:"no-reply@localhost"`
			Name               string `env:"NAME" envDefault:"RCBS"`
		} `envPrefix:"MAIL_"`
	} `envPrefix:"RCBS_"`
}

func Get() config {
	return cfg
}
