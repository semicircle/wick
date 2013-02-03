package main

type Token struct {
    text string
    weight int
    src []Messager
}

func NewToken() Tokener {
    return &Token{"", 0, make([]Messager, 0)}
}

func (t *Token) Text() string {
    return t.text
}

func (t *Token) SetText(s string) error {
    t.text = s
    return nil
}

func (t *Token) Weight() int {
    return t.weight
}

func (t *Token) SetWeight(w int) error {
    t.weight = w
    return nil
}

func (t *Token) SrcMsg() []Messager {
    return t.src
}

func (t *Token) SetSrcMsg(m Messager) error{
    t.src = append(t.src, m)
    return nil
}


