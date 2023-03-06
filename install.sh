if [ -d "/usr/bin/comx" ]; then
    sudo rm /usr/bin/comx
fi

cd /usr/bin
sudo wget https://github.com/jareer12/comx/releases/download/Stable/comx-linux64 -o comx
