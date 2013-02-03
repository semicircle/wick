package main

//produce everything.
type Factory struct {
    newtokener func() Tokener "new tokener funcation"
    newtokencontainer func() TokenContainer "new tokencontainer function"
    newmessagerfromjson func(json string) Messager "new messager function"
    newmessager func(string, string, string, string, string) Messager "new messager function"
    newmessagercontainer func() MessagerContainer
}

func (f *Factory) NewTokener(t string, w int, m Messager) Tokener {
    //return &Token{t, w, m}
    tokener := f.newtokener()
    tokener.SetText(t)
    tokener.SetWeight(w)
    if m != nil {
        tokener.SetSrcMsg(m)
    }
    return tokener
}

func (f *Factory) NewTokenContainer() TokenContainer{
    return f.newtokencontainer()
}

func (f *Factory) NewMessagerFromJson(json string) Messager{
    return f.newmessagerfromjson(json)
}

func (f *Factory) NewMessager(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) Messager{
    return f.newmessager(arg1, arg2, arg3, arg4, arg5)
}

func (f *Factory) NewMessagerContainer() MessagerContainer{
    return f.newmessagercontainer()
}

