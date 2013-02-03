package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "strconv"
)

type MmcRecord struct {
    msg Messager
    cnt int
}

type MapMessagerContainer struct {
    c map[string]MmcRecord
}

func NewMapMessagerContainer() MessagerContainer{
    return &MapMessagerContainer{make(map[string]MmcRecord)}
}

//    Active(Messager, int) error
//    Inactive(Messager, int) error
//    Dump(string)
//    Recover(string)
//    Find(string) Messager

func (m *MapMessagerContainer) Active(msg Messager, cnt int) error{
    if v, ok := m.c[msg.Id()]; ok {
        v.cnt += cnt
    } else {
        m.c[msg.Id()] = MmcRecord{msg, cnt}
    }
    return nil
}

func (m *MapMessagerContainer) Inactive(msg Messager, cnt int) error {
    if v, ok := m.c[msg.Id()]; ok {
        v.cnt -= cnt
        if v.cnt == 0 {
            delete(m.c, msg.Id())
        }
    } else {
        fmt.Printf("MapMessagerContainer Inactive failed: %s not found.\n", msg.Id())
    }
    return nil
}

func (m *MapMessagerContainer) Dump(name string) {
    file, err := os.Create(name + ".mmcdump")
    if (err != nil ) {
        return
    }

    defer file.Close()

    encoder := csv.NewWriter(file)

    cnt := 0

    for k,v := range m.c {
        cnt ++
        line := []string{k, v.msg.Text(), v.msg.Url(), v.msg.ImageUrl(), v.msg.Author(), strconv.Itoa(v.cnt)}
        if err := encoder.Write(line); err != nil {
            fmt.Println(err)
        }
    }
    encoder.Flush()
    fmt.Printf("Dump MapMessagerContainer %s successed, %d message saved.\n" , name, cnt)
}

func (m *MapMessagerContainer) Recover(name string, mf MessagerFactory) {
    file, err := os.Open(name + ".mmcdump")
    if (err != nil) {
        return
    }

    defer file.Close()

    decoder := csv.NewReader(file)

    cnt := 0
    for {
        line, err := decoder.Read()
        if (err != nil) {
            break
        } else {
            msg := mf.NewMessager(line[0], line[1], line[2], line[3], line[4])
            thecnt, _ := strconv.Atoi(line[5])
            m.c[msg.Id()] = MmcRecord{msg, thecnt}
            cnt ++
        }
    }
    fmt.Printf("Recover MapTokenContainer %s successed, %d token recovered.\n", name, cnt)
}

func (m *MapMessagerContainer) Find(idstr string) Messager {
    if v, ok := m.c[idstr]; ok {
        return v.msg
    }
    return nil
}

func (m *MapMessagerContainer) CountAll() int {
    return len(m.c)
}

func (m *MapMessagerContainer) TokenizeAll(tf TokenFactory, tc TokenContainer) {
    for _,v := range m.c {
        msg := v.msg
        msg.Tokenize(tf, tc, m)
    }
}

