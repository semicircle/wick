package main

import (
//    "fmt"
    "os/exec"
)

func SnapWeibo(path string) {
    cmd := exec.Command("ruby", "snap.rb", path)
    cmd.Run()
}

