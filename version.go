package main

import (
    "log"
)

var (
    _GO_VERSION_     = "unknown"
    _OPERATOR_       = "unknown"
    _IP_             = "unknown"
    _GITHUB_VERSION_ = "unknown"
    _DATE_           = "unknown"
    _SYSTEM_         = "unknown"
)

func ShowVersion() {
    log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
    log.Println("go_version:", _GO_VERSION_)
    log.Println("operator:", _OPERATOR_)
    log.Println("ip:", _IP_)
    log.Println("github_flag:", _GITHUB_VERSION_)
    log.Println("date:", _DATE_)
    log.Println("system:", _SYSTEM_)
    log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}