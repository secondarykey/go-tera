package main

import (
	"./api"
	. "./config"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {

	// flags
	err := Setting()
	if err != nil {
		log.Println(err)
		return
	}

	err = Load("tera.ini")
	if err != nil {
		log.Println(err)
		return
	}

	err = os.MkdirAll(Config.Data.Path, 0777)
	if err != nil {
		log.Println(err)
		return
	}

	err = start()
	if err != nil {
		log.Println(err)
		return
	}
}

func start() error {
	return makeTree("/")
}

func makeTree(name string) error {
	log.Println(name)
	files, err := api.Listdir(name)
	if err != nil {
		return err
	}

	log.Println("OK!")
	log.Println(len(files))

	for _, elm := range files {
		realPath := api.RequestName(elm.Href)
		realPath, _ = url.QueryUnescape(realPath)
		if realPath == "" || realPath == name || realPath == ".zfs/" {
			continue
		}

		log.Println(realPath)
		if elm.IsDir {
			err = os.MkdirAll(Config.Data.Path+"/"+realPath, 0777)
			if err != nil {
				return err
			}
			err = makeTree(realPath)
			if err != nil {
				return err
			}
		} else {
			data, err := api.Get(realPath)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(Config.Data.Path+"/"+realPath, data, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
