package gutil

import (
	"github.com/carsonsx/log4g"
	"io/ioutil"
	"encoding/json"
)

func ListenJsonFile(v interface{}, notifyFunc func(filepath string, v interface{}) error, filenames ...string) bool {
	log4g.Info("try to listen json files: %v", filenames)
	loadOne := false
	for _, filename := range filenames {
		if ListenFileChanged(filename, v, loadJsonFile, notifyFunc) {
			loadOne = true
		}
	}
	return loadOne
}

func LoadJsonFile(v interface{}, notifyFunc func(filepath string, v interface{}) error, filepaths ...string) bool {
	log4g.Info("try to load json files: %v", filepaths)
	loadOne := false
	for _, filepath := range filepaths {
		if loadJsonFile(filepath, v) != nil {
			if notifyFunc != nil {
				if notifyFunc(filepath, v) == nil {
					loadOne = true
				}
			}
		}
	}
	return loadOne
}

func loadJsonFile(filename string, v interface{}) error {

	log4g.Debug("try to load json data file: %s", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log4g.Debug(err)
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		log4g.Debug(err)
		return err
	}
	log4g.Info("loaded json data file: %s", filename)
	log4g.Info(log4g.JsonString(v))
	return nil
}