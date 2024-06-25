#!/usr/bin/env python3
'Build script'

import os
import utils

CMD = utils.get_cmd()

BIN_BUILD_PATH = os.path.normpath('build/bin')

utils.make_dir(BIN_BUILD_PATH)


CMD_PATH = os.path.join('.', os.path.normpath(f"./cmd/{CMD}"))
BIN_PATH = os.path.join('.', BIN_BUILD_PATH, f"{CMD}{utils.BIN_EXT}")

utils.run(f'go build -o {BIN_PATH} {CMD_PATH}')
