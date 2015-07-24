#!/bin/bash

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/mickep76/tf"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
cp -r * ${SRCDIR}
(
    echo ${GOPATH}
    cd ${SRCDIR}
    go install .
)
