## Variables
DIR=bin
SUBDIR=normal
NAME=comx

## Create the target directory if does't exist
if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

if [ ! -d "./$DIR/$SUBDIR" ]; then
    mkdir "./$DIR/$SUBDIR"
fi

## Build For Windows
set GOOS=windows
set GOHOSTOS=windows
go build -o ./$DIR/$SUBDIR/$NAME-win64.exe
## Build For Linux Distro(s)
set GOOS=linux
set GOHOSTOS=linux
go build -o ./$DIR/$SUBDIR/$NAME-linux64

## Build for Mac
set GOOS=darwin
setGOHOSTOS=darwin
go build -o ./$DIR/$SUBDIR/$NAME-mac64
