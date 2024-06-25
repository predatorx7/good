#!/usr/bin/env python3
'Run script'

import os
import utils

CMD=utils.get_cmd()

PYTHON_CMD = utils.for_nt_or('python', 'python3')

utils.run(f"{PYTHON_CMD} {os.path.normpath('./scripts/build.py')} {CMD}")

utils.run(os.path.normpath(f"./build/bin/{CMD}{utils.BIN_EXT}"))
