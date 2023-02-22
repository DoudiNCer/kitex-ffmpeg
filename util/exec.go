package util

import (
	"bytes"
	"os"
	"os/exec"
)

// 执行 ffmpeg
func ExecFfmpeg(workPath string, args []string) (sout, serr string, err error) {
	err = os.Chdir(workPath)
	if err != nil {
		return "", "", err
	}
	cmd := exec.Command("ffmpeg", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err = cmd.Run()
	sout, serr = string(stdout.Bytes()), string(stderr.Bytes())
	return
}
