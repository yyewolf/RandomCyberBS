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
}

func Get() config {
	return cfg
}
