package util

import (
	"fmt"
	"time"
)

func MakePath(file string, fileType string) string {
	now := toString(time.Now().Unix())
	path := getPath(fileType)
	path = path + now + file
	return path
}

func toString(now int64) string {
	return fmt.Sprint(now)
}
