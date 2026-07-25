# Build for Linux
$env:GOOS = "linux"
$env:GOARCH = "amd64"

$deployIP = "192.168.18.137"  # Test-server
#$deployIP = "192.168.18.126" # Minecraft

go build -C ./gameserverManager -o ../gameserver
go build -C ./mockServer -o ../mockServerTest

if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed."
    exit 1
}

# Deploy to server - app
scp .\gameserver root@[$deployIP]:/opt/gameserver

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed - .\gameserver"
    exit 1
}

# Deploy to server - env
scp .\mgr.env root@[$deployIP]:/opt/mgr.env

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed - \mgr.env"
    exit 1
}

# Deploy to server - mockServer
scp .\mockServerTest root@[$deployIP]:/opt/mockServerTest

if ($LASTEXITCODE -ne 0) {
    Write-Host "Copy failed - .\mockServerTest"
    exit 1
}

# CHMOD both files
ssh "root@$deployIP" "chmod +x /opt/gameserver; chmod +x /opt/mockServerTest"

if ($LASTEXITCODE -ne 0) {
    Write-Host "SSH Failed - CHMOD."
    exit 1
}

Write-Host "Deployment successful!"