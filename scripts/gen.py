#!/usr/bin/env python3
'Generate script'

import utils

CMD = utils.get_cmd()

utils.make_dir(f'pkg/{CMD}')

utils.write_in_file(f'pkg/{CMD}/{CMD}.go', f"""package {CMD}

func Solve(input string) {utils.OPEN_CURLY_CHAR}
    return ""
{utils.CLOSE_CURLY_CHAR}
""")


utils.make_dir(f'cmd/{CMD}')

utils.write_in_file(f'cmd/{CMD}/main.go', f"""package main

import (
    "fmt"
    "good/pkg/{CMD}"
)

func trySolution(input string) {utils.OPEN_CURLY_CHAR}
	result := {CMD}.Solve(input)
	fmt.Printf("result: %s", result)
{utils.CLOSE_CURLY_CHAR}

func main() {utils.OPEN_CURLY_CHAR}
	trySolution("")
{utils.CLOSE_CURLY_CHAR}
""")
