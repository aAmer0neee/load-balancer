package domain

type Cfg struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
		Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	} `yaml:"server" env-required:"true"`

	Services struct {
		Pool []string `yaml:"pool" env-required:"true"`
	} `yaml:"services"`

	Health struct {
		Timeout int `yaml:"timeout_ms" env-default:"500"`
		Ticker  int `yaml:"ticker_ms" env-default:"5000"`
	} `yaml:"health"`

	Limiter struct {
		Capacity uint32 `yaml:"capacity" env-default:"500"`
		Ticker   int    `yaml:"ticker_ms" env-default:"2000"`
	} `yaml:"limiter"`
}
