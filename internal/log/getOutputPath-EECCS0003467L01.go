package log

import (
	"GOKIT_v001/internal/conf"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getOutputPath(key string) []string {
	list := conf.Of(key)
	mode := list["mode"]
	var retList []string
	if strings.Contains(mode, "STDOUT") {
		retList = append(retList, "stdout")
	}
	if strings.Contains(mode, "STDERR") {
		retList = append(retList, "stderr")
	}
	if strings.Contains(mode, "FILE") {
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)

		filename := fmt.Sprintf("\"%v\\%v%v-%v.log\"", basepath, list["filepath"], list["file"], time.Now().Format("20060102"))
		retList = append(retList, filename)
	}

	return retList
}
