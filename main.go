package main

import (
	"github.com/DodiNCer/kitex-ffmpeg/env"
	kitex_ffmpeg "github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg/kitexffmpeg"
	"log"
)

func main() {
	env.Init()
	svr := kitex_ffmpeg.NewServer(new(KitexFfmpegImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
