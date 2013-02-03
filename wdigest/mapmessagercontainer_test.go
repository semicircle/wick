package main

import (
    "testing"
)

func TestDumpMMC(t *testing.T) {
    f := getFactory("simple")
    c := f.NewMessagerContainer()
    msg1 := f.NewMessager("id1", "text", "http://234.234123",  "http://imageurl.com/123123", "author")
    msg2 := f.NewMessager("id2", "text", "http://234.234123",  "http://imageurl.com/123123", "author")
    msg3 := f.NewMessager("id3", "text", "http://234.234123",  "http://imageurl.com/123123", "author")
    msg4 := f.NewMessager("id4", "text", "http://234.234123",  "http://imageurl.com/123123", "author")
    msg5 := f.NewMessager("id5", "text", "http://234.234123",  "http://imageurl.com/123123", "author")

    c.Active(msg1, 10)
    c.Active(msg2, 11)
    c.Active(msg3, 12)
    c.Active(msg4, 13)
    c.Active(msg5, 14)

    c.Dump("TestDumpMMC")

    c2 := f.NewMessagerContainer()
    c2.Recover("TestDumpMMC",f)

    if msgX := c2.Find("id4"); msgX == nil {
        t.Errorf("Recover Failed: Find")
    }
}
