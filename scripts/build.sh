## Variables
DIR=bin/normal
NAME=comx

## Create the target directory if does't exist
if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

## Build For Windows
set GOOS=windows
set GOHOSTOS=windows
go build -o ./$DIR/$NAME-win64.exe
## Build For Linux Distro(s)
set GOOS=linux
set GOHOSTOS=linux
go build -o ./$DIR/$NAME-linux64

## Build for Mac
set GOOS=darwin
setGOHOSTOS=darwin
go build -o ./$DIR/$NAME-mac64
