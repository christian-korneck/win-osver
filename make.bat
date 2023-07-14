@echo off
setlocal

go install github.com/akavel/rsrc@latest
rsrc.exe -manifest osver.manifest -o rsrc.syso
go build .