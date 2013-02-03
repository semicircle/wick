package main

//tokener, a thing that can be used as a token.
//what's this mean... I can think too clear to start coding, so let it be.
type Tokener interface {
    SetText(string) error
    SetWeight(int) error
    SetSrcMsg(Messager) error
    Text() string
    Weight() int
    SrcMsg() []Messager
}

//which generate Token
type Tokenizer interface {
    Tokenize(TokenFactory, TokenContainer, MessagerContainer) error
}

//Data backend for Token
type TokenContainer interface {
    Active(Tokener) error
    Inactive(Tokener) error
    Search(string) Tokener
    All() []Tokener
    CountAll() int
    Dream(advance int, forget int, c TokenContainer) (int, int)
    Fade(times int)
    Dump(name string)
    Recover(name string, tf TokenFactory)
}

//which produce token
type TokenFactory interface {
    NewTokener(t string, w int, m Messager) Tokener
}

//which produce/provide TokenContainer
type TokenContainerFactory interface {
    NewTokenContainer() TokenContainer
}


