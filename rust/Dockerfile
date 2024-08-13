FROM 117503445/dev-base
LABEL maintainer="117503445"
RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman -Syu rustup --noconfirm
RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
COPY rust/config.toml /root/.cargo/config.toml
RUN rustup default stable
RUN cargo install cargo-watch
ENV VSC_EXTS="rust-lang.rust-analyzer,${VSC_EXTS}"
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"