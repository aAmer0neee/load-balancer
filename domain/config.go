package domain

type Cfg struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	} `yaml:"server" env-required:"true"`

	Services struct {
		Pool []string `yaml:"pool" env-required:"true"`
	} `yaml:"services"`
}
