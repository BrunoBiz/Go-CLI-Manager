# Build for Linux
$env:GOOS = "linux"
$env:GOARCH = "amd64"

go build -o gameserver

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed."
    exit 1
}

# Deploy to server
scp .\gameserver root@192.168.18.126:/tmp/gameserver
scp .\mgr.env root@192.168.18.126:/tmp/mgr.env

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed."
    exit 1
}

Write-Host "Deployment successful!"