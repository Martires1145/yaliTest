package util

import (
	"cmdTest/internal/dto/model"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

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
	defer src.Close().Error()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close().Error()

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

func MakeModelPath(params *model.ParamsJson) string {
	var paramsExtra *model.ParamsExtra
	if params.UseExtra {
		paramsExtra = params.PE
	} else {
		paramsExtra = &model.DefaultParams
	}
	path := fmt.Sprintf("%s_%s_%s_%s_ft%s_sl%s_ll%s_pl%s_dm%s_nh%s_el%s_dl%s_df%s_fc%s_eb%s_dt%s_sc%s_op%s_%s_%s",
		params.PU.TaskName,
		params.PU.ModelID,
		params.PU.Model,
		params.PU.Data,
		params.PU.Features,
		params.PU.SeqLen,
		params.PU.LabelLen,
		params.PU.PredLen,
		paramsExtra.DModel,
		paramsExtra.NHeads,
		params.PU.ELayers,
		params.PU.DLayers,
		paramsExtra.DFF,
		params.PU.Factor,
		paramsExtra.Embed,
		paramsExtra.Distil,
		params.PU.Scale,
		params.PU.Optim,
		params.PU.Des,
		"0",
	)
	return path
}
