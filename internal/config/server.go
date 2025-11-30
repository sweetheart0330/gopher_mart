package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Options struct {
	Host        string `env:"RUN_ADDRESS"`
	URIDatabase string `env:"DATABASE_URI"`
	AccrualAddr string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	AuthPepper  string `env:"AUTH_PEPPER"`
	AuthCost    int    `env:"AUTH_COST"`
}

func GetOptions() (host Options, err error) {
	fl := Options{}
	flag.StringVar(&fl.Host, "a", ":8080", "address and port to send requests")
	flag.StringVar(&fl.URIDatabase, "d", "", "database uri")
	flag.StringVar(&fl.AccrualAddr, "r", "", "accrual system address")
	flag.StringVar(&fl.AuthPepper, "p", "", "auth pepper")
	flag.IntVar(&fl.AuthCost, "c", 8, "auth cost")

	flag.Parse()

	err = env.Parse(&fl)
	if err != nil {
		return Options{}, fmt.Errorf("failed to parse server flags, err: %w", err)
	}

	return fl, nil
}
