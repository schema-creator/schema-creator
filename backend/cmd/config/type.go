package config

var Config *config

type Env string

const (
	EnvProduction  Env = "production"
	EnvDevelopment Env = "development"
	ENVLocal       Env = "local"
	EnvTesting     Env = "testing"
)

type config struct {
	App struct {
		Addr        string `env:"SERVER_ADDR" envDefault:":8080"`
		Env         Env    `env:"ENV"`
		AllowOrigin string `env:"ALLOW_ORIGINS"`
	}

	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		Name     string `env:"DATABASE_NAME"`
	}
}
