package support

import (
	"os"
	"strings"
)

func CreateDirFile(path string) (*os.File, error) {
	res := strings.Split(path, "/")
	res = res[0 : len(res)-1]
	for i := 1; i <= len(res); i++ {
		dir := strings.Join(res[0:i], "/")
		if !PathExists(dir) {
			os.Mkdir(dir, 0777)
		}
	}
	return os.Create(path)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
