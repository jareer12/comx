## Build compressed and optimized for portable use.
## Variables
DIR=bin
SUBDIR=portable
FULLDIR=$DIR/$SUBDIR
NAME=comx
LD_FLAGS='-s -w'
UPX_CMDS='--lzma --best'

## Create the target directory if does't exist
if [ ! -d "./$DIR" ]; then
    mkdir "./$DIR"
fi

if [ ! -d "./$FULLDIR" ]; then
    mkdir "./$FULLDIR"
fi

## Build For Windows
GOOS=windows GOHOSTOS=windows go build -o ./$FULLDIR/$NAME-win64-portable.exe -ldflags="$LD_FLAGS"
upx ./$FULLDIR/$NAME-win64-portable.exe $UPX_CMDS

## Build For Linux Distro(s)
GOOS=linux GOHOSTOS=linux go build -o ./$FULLDIR/$NAME-linux64-portable -ldflags="$LD_FLAGS"
upx ./$FULLDIR/$NAME-linux64-portable $UPX_CMDS

## Build for Mac
GOOS=darwin GOHOSTOS=darwin go build -o ./$FULLDIR/$NAME-mac64-portable -ldflags="$LD_FLAGS"
upx ./$FULLDIR/$NAME-mac64-portable $UPX_CMDS
