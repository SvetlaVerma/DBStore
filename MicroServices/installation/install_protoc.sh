#!/bin/bash

set -e

sudo chmod +x bin/protoc
sudo mv bin/protoc /usr/bin/protoc
cd ${CURDIR}
