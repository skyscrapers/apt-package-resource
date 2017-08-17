#!/bin/bash

set -e -x

# The code is located in /apt-package-resource
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./apt-package-resource to $GOPATH/src/github.com/skyscrapers/apt-package-resource
mkdir -p src/github.com/skyscrapers/
cp -R ./hello-go src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/skyscrapers/apt-package-resource
ls -lat

make
