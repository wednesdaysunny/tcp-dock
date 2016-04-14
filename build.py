#!/usr/bin/env python
from __future__ import print_function
import os
import time
import sys

TARGET_EXE_PATH = "./bin/go_tcp"
SYSTEM_FLAG = "linux"

class BuildProcess(object):
    def __init__(self, version, operator, ip, data, system, git_version):
        self.__version = version
        self.__operator = operator
        self.__ip = ip
        self.__data = data
        self.__system = system
        self.__git_version = git_version

        self.__ldflags = '-ldflags "%s %s %s %s %s %s"' \
            %(self.__version, self.__operator, self.__ip, self.__data, self.__system, self.__git_version)

    def process_build(self):
        target = "" if not TARGET_EXE_PATH else "-o " + TARGET_EXE_PATH
        if SYSTEM_FLAG == "windows":
            if target[-4:] != ".exe":
                target += ".exe"
        build_cmd = 'CGO_ENABLED=0 GOOS=%s GOARCH=amd64 go build %s %s' %(SYSTEM_FLAG, self.__ldflags, target)
        print("Start Build", TARGET_EXE_PATH, build_cmd)
        result_content = run_cmd(build_cmd)


def run_cmd(cmd):
    output = os.popen(cmd)
    result_content = output.read()
    output.close()
    result_content = result_content.replace("\n", " ").strip()
    result_content = result_content.replace(" ", "_")
    return result_content


def main():
    global SYSTEM_FLAG
    global TARGET_EXE_PATH
    
    if len(sys.argv) >= 2:
        if sys.argv[1] in ("linux", "windows"):
            SYSTEM_FLAG = sys.argv[1]
        elif len(sys.argv) == 2:
            TARGET_EXE_PATH = sys.argv[-1]
        if len(sys.argv) == 3:
            TARGET_EXE_PATH = sys.argv[-1]

    go_version_string = run_cmd("go version")
    print("Current Go Version:", go_version_string)

    get_ip_cmd = "LC_ALL=C ifconfig  | grep 'inet addr:'| grep -v '127.0.0.1' | cut -d: -f2 | awk '{ print $1}'"
    ip_string = run_cmd(get_ip_cmd)
    print("Current IP:", ip_string)

    user_string = run_cmd("whoami")
    print("Current User:", user_string)

    date_string = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()) 
    date_string = date_string.replace(" ", "_")
    print("Current Time:", date_string)

    system_string = run_cmd("uname -a")
    print("Current System:", system_string)

    git_version = run_cmd("git rev-parse HEAD")
    print("Current GitFlag:", git_version)

    version_content = "-X main._GO_VERSION_='%s'" %(go_version_string)
    operator_content = "-X main._OPERATOR_='%s'" %(user_string)
    ip_content = "-X main._IP_='%s'" %(ip_string)
    date_content = "-X main._DATE_='%s'" %(date_string)

    system_content = "-X main._SYSTEM_='%s'" %(system_string)
    git_version_content = "-X main._GITHUB_VERSION_='%s'" %(git_version)

    build_obj = BuildProcess(version_content, operator_content, ip_content, date_content, system_content, git_version_content)
    build_obj.process_build()

if __name__ == "__main__":
    main()

