NAME=cpm
DIR=target

if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

## Build For Linux Distro(s)
set GOOS=linux set GOHOSTOS=linux go build -o ./$DIR/$NAME-linux64

## Build For Windows
set GOOS=windows set GOHOSTOS=windows go build -o ./$DIR/$NAME-win64.exe

## Build for Mac
set GOOS=darwin set GOHOSTOS=darwin go build -o ./$DIR/$NAME-mac64
