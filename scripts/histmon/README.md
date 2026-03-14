# histmon

```
go install github.com/117503445/histmon@master
```

```
autoload -Uz add-zsh-hook

typeset -g zsh_command_start_time
typeset -g zsh_current_command

preexec() {
    zsh_command_start_time=$(date +%s%3N)  # 毫秒时间戳
    zsh_current_command=$1
}

precmd() {
    local exit_status=$?
    local end_time=$(date +%s%3N)  # 毫秒时间戳
    
    if [[ -n "$zsh_command_start_time" && -n "$zsh_current_command" ]]; then
        # 调用 histmon，重定向所有输出到 /dev/null
        (COMMAND="$zsh_current_command" \
        START_AT="$zsh_command_start_time" \
        END_AT="$end_time" \
        EXIT_STATUS="$exit_status" \
        TOKEN="" \
        ENDPOINT="" \
        histmon >/dev/null 2>&1 &)
        
        # 清理变量
        unset zsh_command_start_time
        unset zsh_current_command
    fi
}
```