package main

import (
	"flag"

	"github.com/berlingoqc/fxbe/api"
	"github.com/berlingoqc/fxbe/auth"
)

func main() {
	var configfile string

	flag.StringVar(&configfile, "config", "", "the configuration file")
	flag.Parse()

	config, err := api.LoadConfig(configfile)
	if err != nil {
		panic(err)
	}

	auth.GetIAuth = func() auth.IAuth {
		d := &auth.DumpAuth{
			Account: make([]*auth.User, 0),
		}
		p, _ := auth.GetSaltedHash("admin")

		d.Account = append(d.Account, &auth.User{
			ID:       0,
			Username: "admin",
			SaltedPW: p,
		})

		return d
	}

	ws, _ := api.GetWebServer(config)
	ws.Start()
}
