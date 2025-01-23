if status is-interactive
    set fish_greeting # Disable greeting

    # set -x all_proxy "socks5://127.0.0.1:1080"; set -x http_proxy $all_proxy; set -x https_proxy $all_proxy

    set -x PATH ~/.local/bin ~/go/bin $PATH
    
    alias dc="docker compose"
    alias dcu="dc up -d"
    alias dcd="dc down"
    alias dcl="dc logs -f"
    alias dcp="dc pull"
    alias dcr="dc restart"
    alias dc-update="dcp && dcu"
    function ta
        tar -cvf $argv[1].tar $argv[1]
    end
    function targz
        tar -zcvf $argv[1].tar.gz $argv[1]
    end
    function untar
        tar -xvf $argv[1]
    end
    function untargz
        tar -zxvf $argv[1]
    end
end