set -ev

curl https://wiki.117503445.top/linux/script/ssh.sh | bash

cat>/etc/timezone<<EOF
Asia/Shanghai
EOF

echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

pacman-key --init

pacman -Sy archlinux-keyring --noconfirm
pacman -Syyu --noconfirm
pacman -S which zsh btop git docker docker-compose cronie nano vim micro net-tools dnsutils inetutils iproute2 traceroute base-devel parted tmux --noconfirm

cat>>/etc/pacman.conf<<EOF
[archlinuxcn]
Server = https://repo.archlinuxcn.org/\$arch
EOF

pacman -Syu archlinuxcn-keyring --noconfirm
pacman -S yay --noconfirm

chsh -s $(which zsh)

sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
git clone --depth=1 https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting
git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions.git ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions

cat>>/etc/environment<<EOF
LANG=en_US.utf-8
LC_ALL=en_US.utf-8
EOF

tee ~/.zshrc <<-'EOF'
DISABLE_UPDATE_PROMPT=true

export ZSH="$HOME/.oh-my-zsh"

ZSH_THEME="eastwood"

ENABLE_CORRECTION="false"
DISABLE_AUTO_TITLE="true"

plugins=(
    git
    zsh-autosuggestions
    zsh-syntax-highlighting
    sudo
    extract
)

# export http_proxy=http://127.0.0.1:1080 && export https_proxy=http://127.0.0.1:1080

source $ZSH/oh-my-zsh.sh

export PATH=/opt/miniconda3/bin:~/.local/bin:~/go/bin:$PATH
export GOPATH=$HOME/go

# export TERMINFO=/usr/share/terminfo # fix conda

[[ "$TERM_PROGRAM" == "vscode" ]] && . "$(code --locate-shell-integration-path zsh)"

# create .tar
ta() { tar -cvf $1.tar $1; }
# create .tar.gz
targz() { tar -zcvf $1.tar.gz $1; }
# extract .tar
untar() { tar -xvf $1; }
# extract .tar.gz
untargz() { tar -zxvf $1; }

alias "dc-update"="docker compose pull && docker compose up -d"
EOF

