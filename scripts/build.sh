NAME=cpm
DIR=bin
LD_FLAGS='-s'

## Create the target directory if does't exist
if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

## Build For Windows
set GOOS=windows
set GOHOSTOS=windows
go build -o ./$DIR/$NAME-win64.exe -ldflags="$LD_FLAGS"

## Build For Linux Distro(s)
set GOOS=linux
set GOHOSTOS=linux
go build -o ./$DIR/$NAME-linux64 -ldflags="$LD_FLAGS"

## Build for Mac
set GOOS=darwin
setGOHOSTOS=darwin
go build -o ./$DIR/$NAME-mac64 -ldflags="$LD_FLAGS"
