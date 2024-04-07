

# RUN /scripts/install_vsc_ext.py nvarner.typst-lsp tomoki1207.pdf
s = input()
# s = 'RUN /scripts/install_vsc_ext.py nvarner.typst-lsp tomoki1207.pdf'

# get ext list
s = s[s.find('.py')+4:]

exts = s.split()

# ENV VSC_EXTS="ms-vscode.cpptools,ms-vscode.cmake-tools,${VSC_EXTS}"

ext = ','.join(exts)

ss = f'ENV VSC_EXTS="{ext},' + '${VSC_EXTS}"'

print(ss)