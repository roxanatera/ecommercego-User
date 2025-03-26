@echo off
git add .
git commit -m "update"
git push

set GOOS=linux
set GOARCH=amd64
go build -o bootstrap main.go

:: Eliminar archivos anteriores
del main.zip 2> nul
del bootstrap 2> nul

:: Crear nuevo ZIP (solo el binario)
powershell -Command "Compress-Archive -Path .\bootstrap -DestinationPath .\main.zip -CompressionLevel Optimal"

:: Limpiar binario temporal
del bootstrap 2> nul

echo ¡Despliegue generado en main.zip!