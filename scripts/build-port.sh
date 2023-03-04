## Build compressed and optimized for portable use.
## Variables
DIR=bin/portable
NAME=comx
LD_FLAGS='-s -w'
UPX_CMDS='--lzma --best'

## Create the target directory if does't exist
if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

## Build For Windows
set GOOS=windows
set GOHOSTOS=windows
go build -o ./$DIR/$NAME-win64-portable.exe -ldflags="$LD_FLAGS"
upx ./$DIR/$NAME-win64 $UPX_CMDS

## Build For Linux Distro(s)
set GOOS=linux
set GOHOSTOS=linux
go build -o ./$DIR/$NAME-linux64-portable -ldflags="$LD_FLAGS"
upx ./$DIR/$NAME-linux64 $UPX_CMDS

## Build for Mac
set GOOS=darwin
setGOHOSTOS=darwin
go build -o ./$DIR/$NAME-mac64-portable -ldflags="$LD_FLAGS"
upx ./$DIR/$NAME-mac64 $UPX_CMDS
