package main

import (
	"os"

	"github.com/berlingoqc/fxbe/api"
	"github.com/berlingoqc/fxbe/auth"
	"github.com/berlingoqc/fxbe/files"
)

func main() {

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

	api.Context[files.RootKey] = os.Getenv("HOME")
	api.StartWebServer()
}
