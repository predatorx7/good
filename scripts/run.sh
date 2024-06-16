#!/usr/bin/env bash

CMD=$1

./scripts/build.sh $CMD;

./build/bin/$CMD;
