package config

import (
	"github.com/BurntSushi/toml"
)

type setting struct {
	Auth auth
	Data data
}

type auth struct {
	Username string
	Password string
	Url      string
}

type data struct {
	Path string
}

var Config setting

func Load(file string) error {
	_, err := toml.DecodeFile(file, &Config)
	if err != nil {
		return err
	}
	return nil
}
