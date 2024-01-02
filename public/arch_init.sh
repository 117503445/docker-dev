set -ev

curl https://wiki.117503445.top/linux/script/ssh.sh | bash

cat>/etc/timezone<<EOF
Asia/Shanghai
EOF

echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
# echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

# https://wiki.archlinux.org/title/Pacman/Package_signing
pacman-key --init
pacman-key --populate

pacman -Sy archlinux-keyring --noconfirm && pacman -Su --noconfirm
pacman -Syu which zsh fish btop git openssh docker docker-compose cronie nano vim micro net-tools dnsutils inetutils iproute2 traceroute base-devel parted tmux --noconfirm

chsh -s /usr/bin/fish

# https://wiki.archlinux.org/title/locale
cat>>/etc/environment<<EOF
LANG=en_US.utf-8
LC_ALL=en_US.utf-8
EOF

mkdir -p ~/.config/fish
cat << EOF > ~/.config/fish/config.fish
if status is-interactive
    set fish_greeting # Disable greeting

    # set -x all_proxy "socks5://127.0.0.1:1080"; set -x http_proxy \$all_proxy; set -x https_proxy \$all_proxy

    set -x PATH ~/.local/bin ~/go/bin \$PATH
    
    alias dc="docker compose"
    alias dcu="dc up -d"
    alias dcd="dc down"
    alias dcl="dc logs -f"
    alias dcp="dc pull"
    alias dcr="dc restart"
    alias dc-update="dcp && dcu"
    function ta
        tar -cvf \$argv[1].tar \$argv[1]
    end
    function targz
        tar -zcvf \$argv[1].tar.gz \$argv[1]
    end
    function untar
        tar -xvf \$argv[1]
    end
    function untargz
        tar -zxvf \$argv[1]
    end
end
EOF

git config --global user.name "117503445"
git config --global user.email t117503445@gmail.com
# https://git-scm.com/docs/git-config#Documentation/git-config.txt-pushdefault
git config --global push.default current # push the current branch to a branch of the same name
git config --global core.editor "code --wait" # VS Code