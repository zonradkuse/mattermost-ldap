package main

import (
	gcfg "gopkg.in/gcfg.v1"

	"log"
)

type config struct {
	Ldap  map[string]string
	Mysql map[string]string
}

func parseConfig(path string) (cfg config) {
	err := gcfg.ReadFileInto(&cfg, path)

	if err != nil {
		log.Fatal(err)
	}

	return
}
