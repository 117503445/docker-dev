#!/usr/bin/env python3

from pathlib import Path
import subprocess
import sys

# usage: ./install_vsc_ext.py <ext1> <ext2> ...
def main():
    shas = list(Path('/root/.vscode-server/bin').glob('*'))
    if len(shas) != 1:
        print(f'len(shas) != 1, shas: {shas}')
        exit(1)

    file_code_server = Path('/root/.vscode-server/bin') / shas[0] / 'bin' / 'code-server'

    for ext in sys.argv[1:]:
        print(f'Installing extension {ext}')
        subprocess.run([file_code_server, '--install-extension', ext])


if __name__ == '__main__':
    main()