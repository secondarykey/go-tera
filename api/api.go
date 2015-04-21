package api

import (
	. "../config"
	"github.com/secondarykey/davgo"
	"log"
)

const (
	BASE = "/dav"
)

var (
	session *davgo.Session
	err     error
)

func getSession() *davgo.Session {

	if session != nil {
		return session
	}
	session, err = davgo.NewSession(Config.Auth.Url)
	if err != nil {
		panic(err)
	}
	session.SetBasicAuth(Config.Auth.Username, Config.Auth.Password)
	return session
}

func Listdir(path string) error {
	dirs, err := getSession().Listdir(path)
	log.Println(dirs)
	for _, elm := range dirs {
		log.Println(elm)
	}
	return err
}
