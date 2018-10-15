package main

import (
	"time"

	"github.com/l-vitaly/goenv"
)

// env name constants
const (
	TESTEnvName  = "TEST"
	TEST2EnvName = "TEST2"
	TEST3EnvName = "TEST3"
)

// Config service configuration
type Config struct {
	TEST  string
	TEST2 int
	TEST3 time.Duration
}

// Get get env config vars
func Get() (*Config, error) {
	cfg := &Config{}
	goenv.StringVar(&cfg.TEST, TESTEnvName, "value for test")
	goenv.IntVar(&cfg.TEST2, TEST2EnvName, 123)
	goenv.DurationVar(&cfg.TEST3, TEST3EnvName, 5*time.Second)

	//goenv.Parse()
	//if cfg.Mongo.URL == "" {
	//	return nil, fmt.Errorf("could not set %s", DBConnStrEnvName)
	//}
	return cfg, nil
}
