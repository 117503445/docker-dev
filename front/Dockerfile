FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu npm yarn pnpm --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN yarn config set registry https://registry.npmmirror.com/
RUN yarn global add typescript
ENV VSC_EXTS="ritwickdey.liveserver,esbenp.prettier-vscode,humao.rest-client,bradlc.vscode-tailwindcss,dsznajder.es7-react-js-snippets,ecmel.vscode-html-css,xabikos.javascriptsnippets,ms-vscode.vscode-typescript-next,hollowtree.vue-snippets,vue.volar,sdras.vue-vscode-snippets,davidbernegger.vuevelocidad,${VSC_EXTS}"