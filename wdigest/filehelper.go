package main

import (
    //"io"
    "io/ioutil"
    "syscall"
    "path/filepath"
    "os"
)

func GetFileCachedJsons(path string) []string {
    lockfd, _ := syscall.Open(filepath.Join(path, "lock"), syscall.O_RDWR | syscall.O_CREAT, 0644)
    syscall.Flock(lockfd, syscall.LOCK_EX)
    defer syscall.Flock(lockfd, syscall.LOCK_UN)
    defer syscall.Close(lockfd)

    fileinfos, _ := ioutil.ReadDir(path)
    ret := make([]string, len(fileinfos) - 2)

    for i, postfile := range fileinfos {
        if postfile.Name() == "latest" || postfile.Name() == "lock" {
            continue
        }

        fullpath := filepath.Join(path, postfile.Name())

        jsonBytes, _ := ioutil.ReadFile(fullpath)
        ret[i] = string(jsonBytes)

        //os.Remove(fullpath)
    }

    return ret
}

func RemoveFiles(path string) {
    lockfd, _ := syscall.Open(filepath.Join(path, "lock"), syscall.O_RDWR | syscall.O_CREAT, 0644)
    syscall.Flock(lockfd, syscall.LOCK_EX)
    defer syscall.Flock(lockfd, syscall.LOCK_UN)
    defer syscall.Close(lockfd)

    fileinfos, _ := ioutil.ReadDir(path)
    for _, postfile := range fileinfos {
        if postfile.Name() == "latest" || postfile.Name() == "lock" {
            continue
        }
        fullpath := filepath.Join(path, postfile.Name())
        os.Remove(fullpath)
    }

    return
}


