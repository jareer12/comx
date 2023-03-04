NAME=cpm

if [ ! -d "./target" ]; then
    mkdir "./target"
fi

## Build For Linux Distro(s)
set GOOS=linux
set GOHOSTOS=linux
go build -o ./target/$NAME-linux64

## Build For Linux Distro(s)
set GOOS=windows
set GOHOSTOS=windows
go build -o ./target/$NAME-win64.exe

## Build for Mac
set GOOS=darwin set GOHOSTOS=darwin go build -o ./target/$NAME-mac64
