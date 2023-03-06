if [ -d "/usr/bin/comx" ]; then
    rm /usr/bin/comx
fi

sudo curl https://github.com/jareer12/comx/releases/download/Stable/comx-linux64 -o comx
