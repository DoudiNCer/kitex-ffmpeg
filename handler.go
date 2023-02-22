package main

import (
	"context"
	kitex_ffmpeg "github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg"
	"github.com/DodiNCer/kitex-ffmpeg/util"
)

// KitexFfmpegImpl implements the last service interface defined in the IDL.
type KitexFfmpegImpl struct{}

// InitWorkspace implements the KitexFfmpegImpl interface.
func (s *KitexFfmpegImpl) InitWorkspace(ctx context.Context, req *kitex_ffmpeg.InitWorkspaceRequest) (resp *kitex_ffmpeg.InitWorkspaceResponse, err error) {
	token, err := util.InitWorkEnv(req.Whoami)
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return
}

// DropWorkspace implements the KitexFfmpegImpl interface.
func (s *KitexFfmpegImpl) DropWorkspace(ctx context.Context, req *kitex_ffmpeg.DropWorkspaceRequest) (resp *kitex_ffmpeg.DropWorkspaceResponse, err error) {
	err = util.DropWorkDir(req.Whoami, req.Token)
	if err != nil {
		return nil, err
	}
	return
}
