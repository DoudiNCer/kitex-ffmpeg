package main

import (
	"github.com/DodiNCer/kitex-ffmpeg/env"
	kitex_ffmpeg "github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg/kitexffmpeg"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

const ADDR = "0.0.0.0:9427"

func main() {
	env.Init()
	addr, _ := net.ResolveTCPAddr("tcp", ADDR)
	svr := kitex_ffmpeg.NewServer(new(KitexFfmpegImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
