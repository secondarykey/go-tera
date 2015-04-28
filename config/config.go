package config

import (
	"encoding/gob"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/user"
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

func init() {
	gob.Register(&setting{})
}

func Setting() error {

	u, err := user.Current()
	if err != nil {
		return err
	}
	log.Println(u.HomeDir)

	//setting file exists

	//user

	//pass

	//url

	//Get Session

	//data dir

	return nil
}

func Load(file string) error {
	_, err := toml.DecodeFile(file, &Config)
	if err != nil {
		return err
	}

	return encode(&Config)
}

func encode(data *setting) error {
	// create a file
	dataFile, err := os.Create(".go-tera.gob")
	if err != nil {
		return err
	}
	defer dataFile.Close()

	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(data)
	return nil
}
