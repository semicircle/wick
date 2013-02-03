package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "strconv"
)

type MapTokenContainer struct {
    m map[string] Tokener
}

func NewMapTokenContainer() TokenContainer {
    return &MapTokenContainer{make(map[string]Tokener)}
}

func (c *MapTokenContainer) Active(a Tokener) error {
    if t, ok := c.m[a.Text()]; ok {
        t.SetWeight(a.Weight() + t.Weight())
        for _, msg := range a.SrcMsg() {
            t.SetSrcMsg(msg)
        }
        //fmt.Printf("a: %d, t: %d", a.Weight(), t.Weight())
    } else {
        c.m[a.Text()] = a
    }
    return nil
}

func (c *MapTokenContainer) Inactive(a Tokener) error {
    return nil
}

func (c *MapTokenContainer) Search(text string) Tokener {
    if t, ok := c.m[text]; ok {
        return t
    }
    return nil
}

func (c *MapTokenContainer) All() []Tokener {
    tokenslice := make([]Tokener, 0, len(c.m))
    for _,v := range c.m {
        tokenslice = append(tokenslice, v)
    }
    return tokenslice
}

func (c *MapTokenContainer) CountAll() int {
    return len(c.m)
}

func (c *MapTokenContainer) Dream(advance int, forget int, advtc TokenContainer) (int, int) {
    countadv, countfgt := 0, 0
    for k,v := range c.m {
        if (v.Weight() < forget) {
            delete(c.m, k)
            countfgt ++
        }
        if (v.Weight() > advance) {
            advtc.Active(v)
            delete(c.m, k)
            countadv ++
        }
    }
    return countadv, countfgt
}

func (c *MapTokenContainer) Fade(times int) {
    for k,v := range c.m {
        v.SetWeight(v.Weight() - times)
        if v.Weight() <= 0 {
            delete(c.m, k)
        }
    }
}

func (c *MapTokenContainer) Dump(name string) {
    file, err := os.Create(name + ".mtcdump")
    if (err != nil ) {
        return
    }

    defer file.Close()

    encoder := csv.NewWriter(file)

    cnt := 0

    for k,v := range c.m {
        cnt ++
        line := []string{k, strconv.Itoa(v.Weight())}
        //fmt.Println("1")
        for _,msg := range v.SrcMsg() {
            if msg == nil {
                fmt.Println("wtf: nil in token.SrcMsg slice !!")
            }
            //fmt.Println(msg.Id())
            line = append(line, msg.Id())
        }
        if err := encoder.Write(line); err != nil {
            fmt.Println(err)
        }
    }

    encoder.Flush()

    fmt.Printf("Dump MapTokenContainer %s successed, %d token saved.\n" , name, cnt)
}

func (c *MapTokenContainer) Recover(name string, tf TokenFactory) {
    file, err := os.Open(name + ".mtcdump")
    if (err != nil) {
        return
    }

    defer file.Close()

    decoder := csv.NewReader(file)

    cnt := 0
    parseloop:
    for {
        line, err := decoder.Read()
        if (err != nil) {
            if _, ok := err.(*csv.ParseError); !ok{
                //fmt.Println(err)
                break parseloop
            }
        }
        weight, _ := strconv.Atoi(line[1])
        t := tf.NewTokener(line[0], weight, nil)
        mcs := getAllMessagerContainer()
        for _, msgid := range line[2:] {
            for _, mc := range mcs {
                if msg:=mc.Find(msgid); msg != nil{
                    t.SetSrcMsg(msg)
                    break
                }
            }
        }
        c.Active(t)
        cnt++
    }
    fmt.Printf("Recover MapTokenContainer %s successed, %d token recovered.\n", name, cnt)
}


