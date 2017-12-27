package gutil

import (
	"github.com/carsonsx/log4g"
	"io/ioutil"
)

func ListenTextFile(notifyFunc func(filepath, content string) error, filepaths ...string) bool {
	log4g.Info("try to listen json files: %v", filepaths)
	loadOne := false
	for _, filepath := range filepaths {
		if ListenFileChanged(filepath, nil, func(filepath string, v interface{}) error {
			conent, err := loadTextFile(filepath)
			if err == nil && notifyFunc != nil {
				return notifyFunc(filepath, conent)
			}
			return err
		}) {
			loadOne = true
		}
	}
	return loadOne
}

func LoadTextFile(notifyFunc func(filepath, content string) error, filepaths ...string) bool {
	log4g.Info("try to load json files: %v", filepaths)
	loadOne := false
	for _, filepath := range filepaths {
		conent, err := loadTextFile(filepath)
		if err == nil && notifyFunc != nil {
			if notifyFunc(filepath, conent) == nil {
				loadOne = true
			}
		}
	}
	return loadOne
}

func loadTextFile(filepath string) (string, error) {
	log4g.Debug("try to load text data file: %s", filepath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log4g.Debug(err)
		return "", err
	}
	content := string(data)
	log4g.Info("loaded text data file: %s", filepath)
	log4g.Info(content)
	return content, nil
}