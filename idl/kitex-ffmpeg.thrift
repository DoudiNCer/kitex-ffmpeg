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

service KitexFfmpeg{
    InitWorkspaceResponse InitWorkspace(1:InitWorkspaceRequest req)
    DropWorkspaceResponse DropWorkspace(1:DropWorkspaceRequest req)
}
