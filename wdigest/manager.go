package main

//come as the most dirty part.
type Manager struct {
    tokencontainer map[string]TokenContainer
    messagercontainer map[string]MessagerContainer
    factories map[string]FactoryInterface
}

//so-called singleton
var manager *Manager

func getFactory(name string) FactoryInterface{
    InitalConfig()
    return manager.factories[name]
}

func getTokenContainer(name string) TokenContainer{
    InitalConfig()
    return manager.tokencontainer[name]
}

func getAllTokenContainer() []TokenContainer{
    InitalConfig()
    s := make([]TokenContainer, 0, len(manager.tokencontainer))
    for _,v := range manager.tokencontainer{
        s = append(s, v)
    }
    return s
}

func getMessagerContainer(name string) MessagerContainer{
    InitalConfig()
    return manager.messagercontainer[name]
}

func getAllMessagerContainer() []MessagerContainer{
    InitalConfig()
    s := make([]MessagerContainer, 0, len(manager.messagercontainer))
    for _, v := range manager.messagercontainer {
        s = append(s, v)
    }
    return s
}

func InitalConfig() {
    if (manager != nil) {
        return
    }

    manager = new(Manager)
    manager.tokencontainer = make(map[string]TokenContainer)
    manager.messagercontainer = make(map[string]MessagerContainer)
    manager.factories = make(map[string]FactoryInterface)

    manager.factories["simple"] = &Factory{NewToken, NewMapTokenContainer, NewWeiboStatusFromJson, NewWeiboStatus, NewMapMessagerContainer}

    //manager.messagercontainer["tc1"] = newMapTokenContainer()

    //manager.tokencontainer["tc1"] = NewMapTokenContainer()
}

func DumpAll() {
    for name, mc := range manager.messagercontainer {
        mc.Dump(name)
    }
    for name, tc := range manager.tokencontainer{
        tc.Dump(name)
    }
}

func RecoverAll() {
    f := getFactory("simple")
    for name, mc := range manager.messagercontainer {
        mc.Recover(name, f)
    }
    for name, tc := range manager.tokencontainer {
        tc.Recover(name, f)
    }
}
