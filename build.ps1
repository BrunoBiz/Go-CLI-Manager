# Build for Linux
$env:GOOS = "linux"
$env:GOARCH = "amd64"

go build -C ./gameserverManager -o ../gameserver

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed."
    exit 1
}

# Deploy to server - app
#scp .\gameserver root@192.168.18.126:/tmp/gameserver
scp .\gameserver root@192.168.18.137:/tmp/gameserver

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed."
    exit 1
}

# Deploy to server - env
#scp .\mgr.env root@192.168.18.126:/tmp/mgr.env
scp .\mgr.env root@192.168.18.137:/tmp/mgr.env

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed."
    exit 1
}

Write-Host "Deployment successful!"