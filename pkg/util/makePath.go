package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

func toString(now int64) string {
	return fmt.Sprint(now)
}

func GetFile(path string) (f []string, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		f = append(f, path+file.Name())
	}

	return
}

func SaveFile(dst string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err

}

func DeleteFile(path string) error {
	err := os.RemoveAll(path)
	if err == nil {
		return nil
	}

	// 如果标准库方法失败，则使用系统命令
	cmd := exec.Command("rm", "-rf", path)
	err = cmd.Run()
	if err != nil {
		return err
	}

	// 确保文件夹被正确删除
	fi, err := os.Lstat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if fi.IsDir() {
			return os.ErrNotExist
		}
	}

	return nil
}
