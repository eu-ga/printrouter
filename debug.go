//package debug provides functionality for printing all routes
package debug

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
)

const MODE = "debug"

func debugPrint(format string, values ...interface{}) {
	fmt.Printf("[debug] "+format, values...)
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

func calculateAbsolutePath(relativePath string) string {
	return joinPaths("/", relativePath)
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
