package confer


import (
	"os"
	"errors"
	"github.com/BurntSushi/toml"
)


func LoadConfig(pathes []string, conf interface{}) (err error){
	if len(os.Args) > 1 {
		pathes = append(os.Args[1:2], pathes...)
	}
	err = errors.New("No path was passed to find the configuration file")
	for _, path := range pathes {
		_, err = toml.DecodeFile(path, conf)
		if err == nil {
			return
		}
	}
	return
}
