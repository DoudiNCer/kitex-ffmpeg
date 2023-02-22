namespace go kitex_ffmpeg

struct InitWorkspaceRequest{
    1: string whoami
}

struct InitWorkspaceResponse{
    1: string token
}
struct DropWorkspaceRequest{
    1: string whoami
    2: string token
}

struct DropWorkspaceResponse{
}

struct File{
    string fileName
    binary content
}

struct UploadFilesRequest{
    1: string token
    2: list<File> files
}

struct RemoteFile{
    string fileName
    string FileID
}

struct UploadFilesResponse{
    1: list<RemoteFile> files
}

struct DownloadFilesRequest{
    1: string token
    2: list<string> fileIDs
}

struct DownloadFilesResponse{
    1: list<File> files
}

struct ExecRequest{
    1: string token
    2: list<string> args
}

struct ExecResponse{
    1: string sout
    2: string serr
}

service KitexFfmpeg{
    InitWorkspaceResponse InitWorkspace(1:InitWorkspaceRequest req)
    DropWorkspaceResponse DropWorkspace(1:DropWorkspaceRequest req)
    UploadFilesResponse UploadFiles(1:UploadFilesRequest req)
    DownloadFilesResponse DownloadFiles(1:DownloadFilesRequest req)
    ExecResponse ExecFfmpeg(1:ExecRequest req)
}
