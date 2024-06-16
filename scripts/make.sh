#!/usr/bin/env bash

CMD=$1

# PKG

mkdir -p pkg/$CMD;

echo "package $CMD

func Solve() {

}
" > pkg/$CMD/$CMD.go;

# CMD

mkdir -p cmd/$CMD;

echo "package main

import (
	\"fmt\"
	\"good/pkg/$CMD\"
)

func main() {
    $CMD.Solve()
}

" > cmd/$CMD/main.go;

