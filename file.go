package gutil

import (
	"time"
	"os"
)

var file_detect_ticker = time.NewTicker(time.Minute)
var detectFiles = make(map[string]*file_changed_info)
var running = false

type file_changed_info struct {
	filepath    string
	value       interface{}
	mod_time    time.Time
	notifyFuncs []func(filepath string, mapping interface{}) error
}

func start() {
	go func() {
		for {
			<- file_detect_ticker.C
			for filename,fileinfo := range detectFiles {
				if fi, err := os.Stat(filename); err == nil {
					if !fileinfo.mod_time.Equal(fi.ModTime()) {
						fileinfo.mod_time = fi.ModTime()
						for _, notifyFunc := range fileinfo.notifyFuncs {
							if notifyFunc != nil {
								notifyFunc(fileinfo.filepath, fileinfo.value)
							}
						}
					}
				}
			}
		}
	}()
}

func ListenFileChanged(filepath string, v interface{}, notifyFuncs ...func(filepath string, v interface{}) error) bool {
	if !running {
		start()
		running = true
	}
	fileinfo := new (file_changed_info)
	fileinfo.filepath = filepath
	fileinfo.value = v
	fileinfo.notifyFuncs = notifyFuncs
	if fi, err := os.Stat(filepath); err == nil {
		fileinfo.mod_time = fi.ModTime()
	} else {
		return false
	}
	detectFiles[filepath] = fileinfo
	for _, notifyFunc := range notifyFuncs {
		if notifyFunc != nil {
			if notifyFunc(filepath, v) != nil {
				return false
			}
		}
	}
	return true
}

func RemoveFileChangedListener(filepath string) {
	delete(detectFiles, filepath)
}
