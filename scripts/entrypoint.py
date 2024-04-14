#!/usr/bin/env python3

from pathlib import Path
import subprocess
import shutil
import os

def main():
    file_lock = Path(__file__).with_suffix('.lock')
    if file_lock.exists():
        pid = file_lock.read_text().strip()
        # check if the process is still running
        is_running = subprocess.run(["kill", "-0", pid], check=True)
        if is_running.returncode == 0:
            print(f"Process {pid} is still running")
            return
        else:
            print(f"Process {pid} is not running, removing lock file")
            file_lock.unlink()
    
    file_lock.write_text(str(os.getpid()))
    
    try:
        subprocess.run(["/vsc_server_setup.py"], check=True)
    except subprocess.CalledProcessError as e:
        print(f"Error: {e}")

    print("=== Update Arch Linux Packages ===")
    subprocess.run(["pacman", "-Syu", "--noconfirm"], check=True)
    print("=== Update Arch Linux Packages Done ===")

    print("~~~ Successfully Init Development Environment, Enjoy! ~~~")


    file_lock.unlink()


    if Path('/entrypoint').exists():
        print("=== Run Custom Entrypoint ===")
        try:
            subprocess.run(["/entrypoint"], check=True)
        except Exception as e:
            print(f"Error: {e}")
        print("=== Run Custom Entrypoint Done ===")


    # if -t 0
    is_tty = os.isatty(0)
    if is_tty:
        subprocess.call(["/bin/fish"])
    else:
        subprocess.call(["tail -f /dev/null"], shell=True)


if __name__ == '__main__':
    main()