#!/bin/sh -e

usage() {
    echo "Usage: build.sh [check|in|out]"
    echo
    exit 2
}


# sanity check -------

if [ $# -ne 1 ]; then
    usage
fi

if [ "$1" != "check" -a "$1" != "in" -a "$1" != "out" ]; then
    usage
fi

TARGET="$1"
XC_OS="${XC_OS:-$(go env GOOS)}"
XC_ARCH="${XC_ARCH:-$(go env GOARCH)}"

# build ------

echo "Building..."

${GOPATH}/bin/gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "pkg/${TARGET}_{{.OS}}_{{.Arch}}/${TARGET}" \
    ./cmd/${TARGET}
