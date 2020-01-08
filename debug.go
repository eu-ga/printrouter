package printrouter

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
)

var mode = "debug"

func debugPrint(format string, values ...interface{}) {
	switch mode {
	case "debug":
		fmt.Printf("["+mode+"] "+format, values...)
	}
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := lastChar(relativePath) == '/' && lastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func calculateAbsolutePath(prefix, relativePath string) string {
	return joinPaths(prefix, relativePath)
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
