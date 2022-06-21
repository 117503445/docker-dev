# check if the network is abroad

# try to connect to google to determine whether user need to use proxy
curl www.google.com -s -o /dev/null --connect-timeout 1
if [ $? == 0 ]
then
    # echo "connect to google.com successed"
    echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
else
    # echo "connect to google.com failed"
    echo "Server = https://mirrors.aliyun.com/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
    export GOPROXY=https://goproxy.cn
fi