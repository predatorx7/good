#!/usr/bin/env bash

mkdir -p build/bin;

CMD=$1

go build -o build/bin/$CMD ./cmd/$CMD;
