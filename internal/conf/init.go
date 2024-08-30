package conf

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"tmfEcho/internal/global"
)

var apiConf []byte
var debugConf []byte

var confMap = make(map[string]string)
var debugConfMap = make(map[string]string)

func init() {
	confMap = make(map[string]string)

	fill(&confMap, apiConf)
	fill(&confMap, read(global.CONF_DEFAULT_FILE))
	fill(&confMap, read(global.CONF_LOCAL_FILE))

	fill(&debugConfMap, debugConf)
	fill(&debugConfMap, read(global.CONF_LOCAL_DEBUG_FILE))

}

func read(fileName string) []byte {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Could not determine the current path:")
		return nil
	}
	currPath := filepath.Dir(file)
	currFile := currPath + fileName
	b, err := os.ReadFile(currFile)

	if err != nil {
		fmt.Println("Could not read file: " + currFile)
		return nil
	}
	return b
}

func fill(confM *map[string]string, file []byte) {
	if file == nil {
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		conf := parse(scanner.Text())
		if len(conf) == 2 {
			k, v := conf[0], conf[1]
			(*confM)[k] = v
		}
	}
}

func parse(line string) []string {
	line = strings.Trim(line, "\t")
	line = strings.Trim(line, " ")
	return strings.SplitN(line, "=", 2)
}
