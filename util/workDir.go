package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/DodiNCer/kitex-ffmpeg/env"
	"os"
)

// 初始化属于用户的工作区
func InitWorkEnv(whoami string) (token string, err error) {
	s256 := sha256.New()
	s256.Write([]byte(whoami))
	stoken := s256.Sum(nil)
	token = hex.EncodeToString(stoken)
	workPath := env.WorkPara + "/" + token
	err = os.Mkdir(workPath, 0600)
	if err != nil {
		return "", err
	}
	return token, err
}

// 检查用户的工作区
func CheckWorkEnv(whoami, token string) (path string, err error) {
	s256 := sha256.New()
	s256.Write([]byte(whoami))
	stoken := s256.Sum(nil)
	ttoken := hex.EncodeToString(stoken)
	if ttoken != token {

		return "", err
	}
	workPath := env.WorkPara + "/" + token
	_, err = os.Stat(workPath)
	if err != nil {
		err = errors.New("kitex-ffmpeg permission error")
		return "", err
	}
	return workPath, nil
}

// 删除用户工作目录
func DropWorkDir(whoami, token string) error {
	workEnv, err := CheckWorkEnv(whoami, token)
	if err != nil {
		return err
	}
	err = os.RemoveAll(workEnv)
	if err != nil {
		return err
	}
	return nil
}
