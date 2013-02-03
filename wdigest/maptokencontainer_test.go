package main

import "testing"

func TestTokenContainer(t *testing.T) {
    var c TokenContainer
    c = getFactory("simple").NewTokenContainer()
    c = c
}

func TestActiveAndSearch(t *testing.T) {
    f := getFactory("simple")
    var c TokenContainer
    c = f.NewTokenContainer() //MapTokenContainer -> TokenContainer
    tokener := f.NewTokener("Node1", 10, nil)
    if err := c.Active(tokener); err != nil {
        t.Errorf("c.Active returns error")
        t.FailNow()
    }
    if tokener2 := c.Search("Node1"); tokener2 == nil {
        t.Errorf("c.Search does NOT find 'Node1'")
        t.FailNow()
    } else if tokener2.Text() != "Node1" {
        t.Errorf("tokener2.Text() != 'Node1'")
    } else if tokener2.Weight() != 10 {
        t.Errorf("tokener2.Weight() != 10")
    }
    tokener4 := f.NewTokener("Node1", 11, nil)
    if err := c.Active(tokener4); err != nil {
        t.Errorf("c.Active returns error while Active the same token twice")
        t.FailNow()
    }
    if tokener3 := c.Search("Node1"); tokener3 == nil {
        t.Errorf("c.Search does NOT find 'Node1'")
        t.FailNow()
    } else if tokener3.Text() != "Node1" {
        t.Errorf("tokener3.Text() != 'Node1'")
    } else if tokener3.Weight() != 21 {
        t.Errorf("tokener3.Weight() != 21")
    }

}

func TestDumpMTC(t *testing.T) {
    f := getFactory("simple")
    var c TokenContainer
    c = f.NewTokenContainer()
    tokener1 := f.NewTokener("Node1", 10, nil)
    tokener2 := f.NewTokener("Node2", 11, nil)
    tokener3 := f.NewTokener("Node3", 12, nil)
    tokener4 := f.NewTokener("Node4", 13, nil)
    tokener5 := f.NewTokener("Node5", 14, nil)

    c.Active(tokener1)
    c.Active(tokener2)
    c.Active(tokener3)
    c.Active(tokener4)
    c.Active(tokener5)

    c.Dump("TestDumpMTC")

    c2 := f.NewTokenContainer()

    c2.Recover("TestDumpMTC", f)

    if c2.CountAll() != 5 {
        t.Errorf("failed on Recover! Counting: %d", c2.CountAll())
        t.FailNow()
    }

    if tok := c2.Search(tokener3.Text()); tok == nil {
        t.Errorf("failed on Recover! Searching")
    } else if tok.Weight() != 12 {
        t.Errorf("failed on Recover! Ezaming Data")
    }

}


