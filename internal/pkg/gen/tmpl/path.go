package tmpl

import "os"

const (
	server  = "./gen/server"
	service = "./gen/service"
	params  = "./gen/params"
	md      = "./gen/postman_md"
)

func isExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
