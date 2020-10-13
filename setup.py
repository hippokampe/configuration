#!/usr/bin/python3
import os
import subprocess
import json
from shutil import which

PATH_CONFIG = '/etc/hippokampe'
PATH_BROWSERS = f'{PATH_CONFIG}/browsers'
PATH_CONFIG_FILENAME = 'general.json'


def is_tool(name):
    return which(name) is not None


def check_dependencies():
    print('Checking dependencies')
    dependencies = ['npm']
    for dependency in dependencies:
        if not is_tool(dependency):
            exit(f'{dependency} is required. Exiting.')


def create_basic():
    if not os.path.exists(PATH_CONFIG):
        os.makedirs(PATH_CONFIG)


def create_config_file(dataRaw: dict):
    filename = f'{PATH_CONFIG}/{PATH_CONFIG_FILENAME}'
    with open(filename, 'w') as outfile:
        json.dump(dataRaw, outfile, indent=4)


def download_browsers():
    print('Downloading browsers')

    cmd = f'sudo PLAYWRIGHT_BROWSERS_PATH={PATH_BROWSERS} npm i -D playwright'
    subprocess.run(cmd.split(' '))

    cmd = f'ls -l {PATH_CONFIG}/browsers'
    result_cmd = subprocess.run(cmd.split(' '), shell=False,
                                capture_output=True)
    out = result_cmd.stdout.decode("utf-8").strip('\n').split('\n')[1:]

    browsers = []
    print('Setting browsers')
    for line in out:
        browser = line.split(' ')[-1]
        name, version = browser.split('-')

        if name == 'chromium':
            path = f'{PATH_BROWSERS}/{browser}/chrome-linux/chrome'
        elif name == 'firefox':
            path = f'{PATH_BROWSERS}/{browser}/firefox/firefox'
        else:
            continue

        print(f'Setting browser: {name}')
        browsers.append({
            'name': name,
            'version': version,
            'path': path
        })

    return browsers


if __name__ == '__main__':
    if os.geteuid() != 0:
        exit("You need to have root privileges to run this script.\nPlease "
             "try again, this time using 'sudo'. Exiting.")

    check_dependencies()
    create_basic()
    browsers = download_browsers()

    data = {
        'browsers': browsers
    }
    create_config_file(data)

    print('Basic config generated successfully')
