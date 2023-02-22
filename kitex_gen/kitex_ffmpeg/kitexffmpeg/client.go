// Code generated by Kitex v0.4.4. DO NOT EDIT.

package kitexffmpeg

import (
	"context"
	kitex_ffmpeg "github.com/DodiNCer/kitex-ffmpeg/kitex_gen/kitex_ffmpeg"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	InitWorkspace(ctx context.Context, req *kitex_ffmpeg.InitWorkspaceRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.InitWorkspaceResponse, err error)
	DropWorkspace(ctx context.Context, req *kitex_ffmpeg.DropWorkspaceRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.DropWorkspaceResponse, err error)
	UploadFiles(ctx context.Context, req *kitex_ffmpeg.UploadFilesRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.UploadFilesResponse, err error)
	DownloadFiles(ctx context.Context, req *kitex_ffmpeg.DownloadFilesRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.DownloadFilesResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kKitexFfmpegClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kKitexFfmpegClient struct {
	*kClient
}

func (p *kKitexFfmpegClient) InitWorkspace(ctx context.Context, req *kitex_ffmpeg.InitWorkspaceRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.InitWorkspaceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InitWorkspace(ctx, req)
}

func (p *kKitexFfmpegClient) DropWorkspace(ctx context.Context, req *kitex_ffmpeg.DropWorkspaceRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.DropWorkspaceResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DropWorkspace(ctx, req)
}

func (p *kKitexFfmpegClient) UploadFiles(ctx context.Context, req *kitex_ffmpeg.UploadFilesRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.UploadFilesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadFiles(ctx, req)
}

func (p *kKitexFfmpegClient) DownloadFiles(ctx context.Context, req *kitex_ffmpeg.DownloadFilesRequest, callOptions ...callopt.Option) (r *kitex_ffmpeg.DownloadFilesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DownloadFiles(ctx, req)
}
