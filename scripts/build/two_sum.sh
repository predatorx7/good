#!/usr/bin/env bash

mkdir -p build/bin;

go build -o build/bin/two_sum ./cmd/two_sum;
