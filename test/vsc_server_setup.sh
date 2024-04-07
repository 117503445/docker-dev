#!/usr/bin/env bash

set -e

URL=https://vsc-server.unidrop.top/
# if URL not end with /, add /
if [[ ! ${URL} =~ /$ ]]; then
    URL="${URL}/"
fi

# check curl or wget

HTTP_TOOL=""
if command -v curl &>/dev/null; then
    HTTP_TOOL="curl"
elif command -v wget &>/dev/null; then
    HTTP_TOOL="wget"
else
    echo "curl or wget is required"
    exit 1
fi

echo "using ${HTTP_TOOL} as HTTP tool, source: ${URL}"

# get latest sha
# get https://webdav.gh.117503445.top:20000/vsc-server/latest-sha

if [ "${HTTP_TOOL}" == "curl" ]; then
    LATEST_SHA=$(curl -s ${URL}latest-sha)
elif [ "${HTTP_TOOL}" == "wget" ]; then
    LATEST_SHA=$(wget -q -O - ${URL}latest-sha)
fi

echo "latest sha: ${LATEST_SHA}"

HOME_DIR=$HOME
if [ -z ${HOME_DIR} ]; then
    USER=$(whoami)
    if [ "${USER}" == "root" ]; then
        HOME_DIR="/root"
    else
        HOME_DIR="/home/${USER}"
    fi
fi

echo "home dir: ${HOME_DIR}"

FILE_NODE="${HOME_DIR}/.vscode-server/bin/${LATEST_SHA}/node"
if [ -f ${FILE_NODE} ]; then
    echo "${FILE_NODE} exists, skip download"
else
    TMP_VSC_DIR="/tmp/vsc-server"
    mkdir -p ${TMP_VSC_DIR}

    VSC_FILE_URL="${URL}${LATEST_SHA}.tar.gz"
    TMP_VSC_FILE="${TMP_VSC_DIR}/${LATEST_SHA}.tar.gz"

    if [ ! -f ${TMP_VSC_FILE} ]; then
        echo "downloading ${VSC_FILE_URL} to ${TMP_VSC_FILE}"
        if [ "${HTTP_TOOL}" == "curl" ]; then
            curl ${VSC_FILE_URL} -o ${TMP_VSC_FILE}
        elif [ "${HTTP_TOOL}" == "wget" ]; then
            wget -O ${TMP_VSC_FILE} ${VSC_FILE_URL}
        fi
        echo "downloaded ${VSC_FILE_URL} to ${TMP_VSC_FILE}"
    fi

    echo "extracting ${TMP_VSC_DIR}/${LATEST_SHA}.tar.gz"
    tar -xzf ${TMP_VSC_DIR}/${LATEST_SHA}.tar.gz -C ${TMP_VSC_DIR}
    echo "extracted ${TMP_VSC_DIR}/${LATEST_SHA}.tar.gz"

    mkdir -p ${HOME_DIR}/.vscode-server/bin/${LATEST_SHA}
    mv ${TMP_VSC_DIR}/vscode-server-linux-x64/* ${HOME_DIR}/.vscode-server/bin/${LATEST_SHA}
fi

# $VSC_EXTS = mhutchie.git-graph,ms-ceintl.vscode-language-pack-zh-hans

# split by comma
IFS=',' read -r -a VSC_EXTS <<<"$VSC_EXTS"
echo "installing extensions: ${VSC_EXTS[@]}"
for ext in "${VSC_EXTS[@]}"; do
    ${HOME_DIR}/.vscode-server/bin/${LATEST_SHA}/bin/code-server --install-extension ${ext}
done