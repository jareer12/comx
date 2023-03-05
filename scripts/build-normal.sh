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
GOOS=windows GOHOSTOS=windows go build -o ./$DIR/$SUBDIR/$NAME-win64.exe

## Build For Linux Distro(s)
GOOS=linux GOHOSTOS=linux go build -o ./$DIR/$SUBDIR/$NAME-linux64

## Build for Mac
GOOS=darwin GOHOSTOS=darwin go build -o ./$DIR/$SUBDIR/$NAME-mac64
