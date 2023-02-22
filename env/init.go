package env

import (
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"os"
)

var WorkPara string

const DEFAULTWORKPARA = "/tmp/kitex_ffmpeg"
const SYSWORKDIRENV = "KITEX_FFMPEG_WORKDIR"

// 用于检查并创建工作目录
func Init() {
	val, ok := os.LookupEnv(SYSWORKDIRENV)
	if ok {
		log.Info("Found Environment Path:", SYSWORKDIRENV, " = ", val, ", checking...")
		WorkPara = val
		os.Setenv(SYSWORKDIRENV, WorkPara)
		err := os.MkdirAll(WorkPara, 0700)
		if err != nil {
			log.Warn("Invalid working directory provided, trying to use default: ", DEFAULTWORKPARA)
			log.Info("Setting Up default workdir: ", DEFAULTWORKPARA)
			WorkPara = DEFAULTWORKPARA
			err := os.MkdirAll(WorkPara, 0700)
			if err != nil {
				panic("Default workdir is invalid too, exiting...")
			}
		}
	} else {
		log.Info("environment variable ", SYSWORKDIRENV, " not found， setting up default workdir: ", DEFAULTWORKPARA)
		WorkPara = DEFAULTWORKPARA
		err := os.MkdirAll(WorkPara, 0700)
		if err != nil {
			panic("Default workdir is invalid, exiting...")
		}
	}
}
