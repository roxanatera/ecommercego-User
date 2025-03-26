git add .
git commit -m "update"
git push
set GOOS=linux
set GOARCH=amd64
go build -o bootstrap main.go
del main.zip
del function.zip
tar.exe -a -cf main.zip main