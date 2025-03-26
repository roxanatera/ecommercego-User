@echo off
git add .
git commit -m "update"
git push

set GOOS=linux
set GOARCH=amd64
go build -o bootstrap main.go 

del main.zip 2> nul
powershell -Command "Compress-Archive -Path .\bootstrap -DestinationPath .\main.zip"