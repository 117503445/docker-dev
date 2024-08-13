FROM 117503445/dev-base
LABEL maintainer="117503445"

# GO
ENV PATH="/root/go/bin:${PATH}"

RUN echo "Server = https://mirrors.kernel.org/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
# CPP
RUN pacman -Syu gcc clang cmake --noconfirm

# FRONT
RUN pacman -Syu npm yarn pnpm --noconfirm

# GO
RUN pacman -Syu go --noconfirm

# JAVA
RUN pacman -Sy jdk-openjdk maven gradle --noconfirm

RUN su - builder -c "yay -Su scala --noconfirm"

# PYTHON
RUN pacman -Syu poetry mypy rye --noconfirm
RUN poetry config virtualenvs.in-project true
RUN su - builder -c "yay -Su miniconda3 --noconfirm"

# RUST
RUN pacman -Syu rustup --noconfirm

RUN echo "Server = https://mirrors.ustc.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist

# FRONT
RUN npm config set registry https://registry.npmmirror.com
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn global add typescript
RUN pnpm config set registry https://registry.npmmirror.com

# GO
WORKDIR /scripts
COPY ./scripts/set_go.sh .
RUN ./set_go.sh && rm set_go.sh
WORKDIR /workspace

# RUST
COPY rust/config.toml /root/.cargo/config.toml
RUN rustup default stable
RUN cargo install cargo-watch
ENV RUSTUP_DIST_SERVER="https://rsproxy.cn"
ENV RUSTUP_UPDATE_ROOT="https://rsproxy.cn/rustup"