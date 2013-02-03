package main

import (
    "fmt"
    //"strings"
    "strconv"
)

func commandLoop() {
    for {
        var cmd,arg1,arg2,arg3,arg4,arg5,arg6 string

        fmt.Printf("#_# ")

        fmt.Scanln(&cmd, &arg1, &arg2, &arg3, &arg4, &arg5, &arg6)

        switch {
        case cmd == "":
            continue
        case cmd == "snapWeibo":
            cmdSnapWeibo(arg1, arg2, arg3)
        case cmd == "clearWeibo":
            cmdClearWeibo(arg1)
        case cmd == "newmc":
            cmdNewMc(arg1, arg2)
        case cmd == "newtc":
            cmdNewTc(arg1, arg2)
        case cmd ==  "exit":
            fmt.Println("Bye")
            goto exit
        case cmd == "dumpall":
            DumpAll()
        case cmd == "recoverall":
            RecoverAll()
        case cmd == "show":
            cmdShow()
        case cmd == "dream":
            cmdDream(arg1, arg2, arg3, arg4)
        case cmd == "fade":
            cmdFade(arg1, arg2)
        case cmd == "tokenize":
            cmdTokenize(arg1, arg2, arg3)

        default:
            fmt.Printf("unknown command: %s", cmd)


        }
    }
    exit:
}

func cmdNewMc( factory string, name string) {
    f := getFactory(factory)
    mc := f.NewMessagerContainer()
    manager.messagercontainer[name] = mc
    fmt.Println("new MessagerContainer ", name, " created.")
}

func cmdNewTc( factory string, name string) {
    f := getFactory(factory)
    mc := f.NewTokenContainer()
    manager.tokencontainer[name] = mc
    fmt.Println("new TokenContainer", name, " created.")
}

func cmdSnapWeibo( factory string, path string, mcname string) {
    f := getFactory(factory)
    mc := getMessagerContainer(mcname)

    jsons := GetFileCachedJsons(path)

    cnt := 0
    for _, json := range jsons {
        msg := f.NewMessagerFromJson(json)
        mc.Active(msg, 0)
        cnt ++
    }
    fmt.Printf("snapWeibo successed, cnt: %d", cnt)
}

func cmdClearWeibo( path string) {
    RemoveFiles(path)
}

func cmdShow() {
    allmc := manager.messagercontainer
    alltc := manager.tokencontainer

    for name, mc := range allmc {
        fmt.Printf("MC [%s] contains %d message\n", name, mc.CountAll())
    }

    for name, tc := range alltc {
        fmt.Printf("TC [%s] contains %d token\n", name, tc.CountAll())
    }
}

func cmdTokenize(factory, mcname, tcname string ) {
    f := getFactory(factory)
    mc := getMessagerContainer(mcname)
    tc := getTokenContainer(tcname)
    mc.TokenizeAll(f, tc)
    fmt.Printf("Tokenize complete, please check 'show'\n")
}

func cmdDream(tcFrom, tcTo, advance, forget string) {
    tc := getTokenContainer(tcFrom)
    tcAdv := getTokenContainer(tcTo)

    adv, _ := strconv.Atoi(advance)
    fgt, _ := strconv.Atoi(forget)

    tc.Dream(adv, fgt, tcAdv)
}

func cmdFade(tcname, strtimes string) {
    times, _ := strconv.Atoi(strtimes)
    tc := getTokenContainer(tcname)
    tc.Fade(times)
}

