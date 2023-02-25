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
	var rep kitex_ffmpeg.InitWorkspaceResponse
	rep.Token = token
	resp = &rep
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

// UploadFiles implements the KitexFfmpegImpl interface.
func (s *KitexFfmpegImpl) UploadFiles(ctx context.Context, req *kitex_ffmpeg.UploadFilesRequest) (resp *kitex_ffmpeg.UploadFilesResponse, err error) {
	workPath, err := util.CheckWorkEnvExist(req.Token)
	var resps kitex_ffmpeg.UploadFilesResponse
	if err != nil {
		return nil, err
	}
	//fileIDs := resps.GetFiles()
	var fileIDs []*kitex_ffmpeg.RemoteFile
	for _, file := range req.Files {
		fileID, err := util.WriteFile(workPath, file)
		if err != nil {
			return nil, err
		}
		var rf kitex_ffmpeg.RemoteFile
		rf.FileID = fileID
		rf.FileName = file.FileName
		fileIDs = append(fileIDs, &rf)
	}
	resps.Files = fileIDs
	return &resps, nil
}

// DownloadFiles implements the KitexFfmpegImpl interface.
func (s *KitexFfmpegImpl) DownloadFiles(ctx context.Context, req *kitex_ffmpeg.DownloadFilesRequest) (resp *kitex_ffmpeg.DownloadFilesResponse, err error) {
	workPath, err := util.CheckWorkEnvExist(req.Token)
	if err != nil {
		return nil, err
	}
	var resps kitex_ffmpeg.DownloadFilesResponse
	//files := resp.GetFiles()
	var files []*kitex_ffmpeg.File
	for _, fileID := range req.FileIDs {
		file, err := util.ReadFile(workPath, fileID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	resps.Files = files
	return &resps, nil
}

// ExecFfmpeg implements the KitexFfmpegImpl interface.
func (s *KitexFfmpegImpl) ExecFfmpeg(ctx context.Context, req *kitex_ffmpeg.ExecRequest) (resp *kitex_ffmpeg.ExecResponse, err error) {
	workPath, err := util.CheckWorkEnvExist(req.Token)
	if err != nil {
		return nil, err
	}
	var rep kitex_ffmpeg.ExecResponse
	sout, serr, err := util.ExecFfmpeg(workPath, req.Args_)

	rep.Sout = sout
	rep.Serr = serr
	resp = &rep
	return resp, err
}
