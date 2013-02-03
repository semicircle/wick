package main

type Messager interface {
    Tokenizer
    Id() string
    Text() string
    Url() string
    ImageUrl() string
    Author() string
}

type MessagerContainer interface {
    Active(Messager, int) error
    Inactive(Messager, int) error
    Dump(string)
    Recover(string, MessagerFactory)
    Find(string) Messager
    CountAll() int
    TokenizeAll(TokenFactory, TokenContainer)
}

type MessagerFactory interface {
    NewMessagerFromJson(string) Messager
    NewMessager(id string, text string, url string, imageurl string, author string) Messager
}

type MessagerContainerFactory interface {
    NewMessagerContainer() MessagerContainer
}
