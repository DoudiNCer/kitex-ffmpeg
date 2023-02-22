package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg"
	"io/ioutil"
	"os"
	"time"
)

const (
	FILENP = "kitex_ffmpeg_file"
	FILENN = "19932023"
)

// 通过文件名计算文件ID（在存储设备上的文件名）并判断是否存在
func getFileID(filename string) (fileID string, err error) {
	md5n := md5.New()
	md5n.Write([]byte(FILENP))
	md5n.Write([]byte(fileID))
	md5n.Write([]byte(time.Now().GoString()))
	md5n.Write([]byte(FILENN))
	fileID = hex.EncodeToString(md5n.Sum(nil))
	_, err = os.Stat(fileID)
	if err == nil {
		err = errors.New("kitex_ffmpeg: file exists")
		return "", err
	}
	return fileID, nil
}

// 将文件写入工作区
func WriteFile(workPath string, file *kitex_ffmpeg.File) (fileID string, err error) {
	id, err := getFileID(file.FileName)
	if err != nil {
		return "", err
	}
	err = os.Chdir(workPath)
	if err != nil {
		return "", err
	}
	nFile, err := os.OpenFile(id, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return "", err
	}
	_, err = nFile.Write(file.Content)
	if err != nil {
		return "", err
	}
	return fileID, nil
}

// 从工作区读取文件
func ReadFile(workPath, fileID string) (file *kitex_ffmpeg.File, err error) {
	err = os.Chdir(workPath)
	if err != nil {
		return nil, err
	}
	nFile, err := os.Open(fileID)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(nFile)
	if err != nil {
		return nil, err
	}
	file.FileName = fileID
	file.Content = data
	return file, nil
}
