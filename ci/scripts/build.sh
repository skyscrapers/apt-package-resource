#!/bin/bash

set -e -x

# The code is located in /apt-package-resource
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD
export PATH=$GOPATH/bin:$PATH

# Now we must move our code from the current directory ./apt-package-resource to $GOPATH/src/github.com/skyscrapers/apt-package-resource
mkdir -p src/github.com/skyscrapers/
cp -R ./apt-package-resource src/github.com/skyscrapers/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/skyscrapers/apt-package-resource
ls -lat

go get -u github.com/golang/dep/cmd/dep
${GOPATH}/bin/dep ensure

make prepare
make
