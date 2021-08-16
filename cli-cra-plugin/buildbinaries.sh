#!/bin/bash

set -e

verify_argument() {
    if [ -z "$2" ]
    then
        printf '{"failed: true, "msg": "%s must be defined"}' "$1"
        exit 1
    fi
}

#increment this before running the script if you plan to publish a new version
VERSION="0.3.2"

BIN_DIR=binaries

verify_argument "VERSION" $VERSION

if [ -d "binaries/" ]; then
    rm -rf binaries/
fi


env GOARC=386 GOOS=linux go build -o ./binaries/doi-cli-linux-386-$VERSION github.ibm.com/oneibmcloud/cli-cra-plugin
LINUX32_CHECKSUM=$(shasum binaries/doi-cli-linux-386-$VERSION | awk '{print $1}')
echo "linux32 binary build! - " $LINUX32_CHECKSUM

env GOARCH=amd64 GOOS=linux go build -o ./binaries/doi-cli-linux-amd64-$VERSION github.ibm.com/oneibmcloud/cli-cra-plugin
LINUX64_CHECKSUM=$(shasum binaries/doi-cli-linux-amd64-$VERSION | awk '{print $1}')
echo "linux64 binary built! - " $LINUX64_CHECKSUM

env GOARCH=386 GOOS=windows go build -o ./binaries/doi-cli-win32-$VERSION.exe github.ibm.com/oneibmcloud/cli-cra-plugin
WIN32_CHECKSUM=$(shasum binaries/doi-cli-win32-$VERSION.exe | awk '{print $1}')
echo "win32 binary built! - " $WIN32_CHECKSUM

env GOARCH=amd64 GOOS=windows go build -o ./binaries/doi-cli-win64-$VERSION.exe github.ibm.com/oneibmcloud/cli-cra-plugin
WIN64_CHECKSUM=$(shasum binaries/doi-cli-win64-$VERSION.exe | awk '{print $1}')
echo "win64 binary built! - " $WIN64_CHECKSUM

env GOARCH=amd64 GOOS=darwin go build -o ./binaries/doi-cli-osx-$VERSION github.ibm.com/oneibmcloud/cli-cra-plugin
CKSUM=$(shasum binaries/doi-cli-osx-$VERSION | awk '{print $1}')
echo "osx binary built! - " $CKSUM
