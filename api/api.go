package api

import (
	. "../config"
	"github.com/secondarykey/davgo"
	"io/ioutil"
	"log"
	"strings"
)

const (
	BASE = "/dav/"
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

func Listdir(path string) ([]davgo.FileInfo, error) {
	dirs, err := getSession().Listdir(path)
	if err != nil {
		log.Println(err)
	}
	return dirs, nil
}

func Get(name string) ([]byte, error) {
	reader, err := getSession().NewReader(name)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(*reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RequestName(name string) string {
	return strings.Replace(name, BASE, "", -1)
}
