'Utilities'

import os
import sys

def run(cmd):
    'Run command'
    print(f'$ {cmd}')
    os.system(cmd)

def get_cmd():
    'Get CMD from argument'
    if len(sys.argv) < 2:
        print('Missing CMD name')
        sys.exit(1)
    return sys.argv[1]

def make_dir(path):
    'Create directory if not exist'
    os.makedirs(name=os.path.abspath(path), exist_ok=True)

def write_in_file(path, content):
    'Write content to file'
    with open(os.path.abspath(path), "w+", encoding='utf8') as f:
        f.write(content)

def for_nt_or(for_nt, for_others):
    'Returns value different for windows and others'
    if os.name == 'nt':
        return for_nt
    return for_others

BIN_EXT = for_nt_or('.exe', '')

OPEN_CURLY_CHAR = '{'
CLOSE_CURLY_CHAR = '}'
