// Code generated by Kitex v0.4.4. DO NOT EDIT.

package kitexffmpeg

import (
	"context"
	kitex_ffmpeg "github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return kitexFfmpegServiceInfo
}

var kitexFfmpegServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "KitexFfmpeg"
	handlerType := (*kitex_ffmpeg.KitexFfmpeg)(nil)
	methods := map[string]kitex.MethodInfo{
		"InitWorkspace": kitex.NewMethodInfo(initWorkspaceHandler, newKitexFfmpegInitWorkspaceArgs, newKitexFfmpegInitWorkspaceResult, false),
		"DropWorkspace": kitex.NewMethodInfo(dropWorkspaceHandler, newKitexFfmpegDropWorkspaceArgs, newKitexFfmpegDropWorkspaceResult, false),
		"UploadFiles":   kitex.NewMethodInfo(uploadFilesHandler, newKitexFfmpegUploadFilesArgs, newKitexFfmpegUploadFilesResult, false),
		"DownloadFiles": kitex.NewMethodInfo(downloadFilesHandler, newKitexFfmpegDownloadFilesArgs, newKitexFfmpegDownloadFilesResult, false),
		"ExecFfmpeg":    kitex.NewMethodInfo(execFfmpegHandler, newKitexFfmpegExecFfmpegArgs, newKitexFfmpegExecFfmpegResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "kitex_ffmpeg",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func initWorkspaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*kitex_ffmpeg.KitexFfmpegInitWorkspaceArgs)
	realResult := result.(*kitex_ffmpeg.KitexFfmpegInitWorkspaceResult)
	success, err := handler.(kitex_ffmpeg.KitexFfmpeg).InitWorkspace(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKitexFfmpegInitWorkspaceArgs() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegInitWorkspaceArgs()
}

func newKitexFfmpegInitWorkspaceResult() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegInitWorkspaceResult()
}

func dropWorkspaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*kitex_ffmpeg.KitexFfmpegDropWorkspaceArgs)
	realResult := result.(*kitex_ffmpeg.KitexFfmpegDropWorkspaceResult)
	success, err := handler.(kitex_ffmpeg.KitexFfmpeg).DropWorkspace(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKitexFfmpegDropWorkspaceArgs() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegDropWorkspaceArgs()
}

func newKitexFfmpegDropWorkspaceResult() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegDropWorkspaceResult()
}

func uploadFilesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*kitex_ffmpeg.KitexFfmpegUploadFilesArgs)
	realResult := result.(*kitex_ffmpeg.KitexFfmpegUploadFilesResult)
	success, err := handler.(kitex_ffmpeg.KitexFfmpeg).UploadFiles(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKitexFfmpegUploadFilesArgs() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegUploadFilesArgs()
}

func newKitexFfmpegUploadFilesResult() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegUploadFilesResult()
}

func downloadFilesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*kitex_ffmpeg.KitexFfmpegDownloadFilesArgs)
	realResult := result.(*kitex_ffmpeg.KitexFfmpegDownloadFilesResult)
	success, err := handler.(kitex_ffmpeg.KitexFfmpeg).DownloadFiles(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKitexFfmpegDownloadFilesArgs() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegDownloadFilesArgs()
}

func newKitexFfmpegDownloadFilesResult() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegDownloadFilesResult()
}

func execFfmpegHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*kitex_ffmpeg.KitexFfmpegExecFfmpegArgs)
	realResult := result.(*kitex_ffmpeg.KitexFfmpegExecFfmpegResult)
	success, err := handler.(kitex_ffmpeg.KitexFfmpeg).ExecFfmpeg(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKitexFfmpegExecFfmpegArgs() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegExecFfmpegArgs()
}

func newKitexFfmpegExecFfmpegResult() interface{} {
	return kitex_ffmpeg.NewKitexFfmpegExecFfmpegResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) InitWorkspace(ctx context.Context, req *kitex_ffmpeg.InitWorkspaceRequest) (r *kitex_ffmpeg.InitWorkspaceResponse, err error) {
	var _args kitex_ffmpeg.KitexFfmpegInitWorkspaceArgs
	_args.Req = req
	var _result kitex_ffmpeg.KitexFfmpegInitWorkspaceResult
	if err = p.c.Call(ctx, "InitWorkspace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DropWorkspace(ctx context.Context, req *kitex_ffmpeg.DropWorkspaceRequest) (r *kitex_ffmpeg.DropWorkspaceResponse, err error) {
	var _args kitex_ffmpeg.KitexFfmpegDropWorkspaceArgs
	_args.Req = req
	var _result kitex_ffmpeg.KitexFfmpegDropWorkspaceResult
	if err = p.c.Call(ctx, "DropWorkspace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadFiles(ctx context.Context, req *kitex_ffmpeg.UploadFilesRequest) (r *kitex_ffmpeg.UploadFilesResponse, err error) {
	var _args kitex_ffmpeg.KitexFfmpegUploadFilesArgs
	_args.Req = req
	var _result kitex_ffmpeg.KitexFfmpegUploadFilesResult
	if err = p.c.Call(ctx, "UploadFiles", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadFiles(ctx context.Context, req *kitex_ffmpeg.DownloadFilesRequest) (r *kitex_ffmpeg.DownloadFilesResponse, err error) {
	var _args kitex_ffmpeg.KitexFfmpegDownloadFilesArgs
	_args.Req = req
	var _result kitex_ffmpeg.KitexFfmpegDownloadFilesResult
	if err = p.c.Call(ctx, "DownloadFiles", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ExecFfmpeg(ctx context.Context, req *kitex_ffmpeg.ExecRequest) (r *kitex_ffmpeg.ExecResponse, err error) {
	var _args kitex_ffmpeg.KitexFfmpegExecFfmpegArgs
	_args.Req = req
	var _result kitex_ffmpeg.KitexFfmpegExecFfmpegResult
	if err = p.c.Call(ctx, "ExecFfmpeg", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
