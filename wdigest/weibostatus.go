package main

import (
    //"fmt"
    "encoding/json"
    "log"
    "strings"
    //"container/list"
    //"io"
    "regexp"
)

//WeiboStatus is both a Messager & Tokenizer
type WeiboStatus struct {
    idstr, text, url, imageUrl string
}

func NewWeiboStatusFromJson(jsonText string) Messager {
    dec := json.NewDecoder(strings.NewReader(jsonText))
    type jsonWeibo struct {
        Text, Url, Bmiddle_pic, Idstr  string
        Retweeted_status *struct {
            Text, Bmiddle_pic string
            User struct {
                Name string
            }
        }
        User struct {
            Name string
        }
    }

    var m jsonWeibo
    if err := dec.Decode(&m); err != nil {
        log.Fatal(err)
    }
    var text, imageUrl string
    if m.Retweeted_status != nil {
        text = m.Text + " //@" + m.Retweeted_status.User.Name + ": " + m.Retweeted_status.Text
        imageUrl = m.Retweeted_status.Bmiddle_pic
    } else {
        text = m.Text
        imageUrl = m.Bmiddle_pic
    }
    return &WeiboStatus{m.Idstr, text, "weibo:" + m.Idstr, imageUrl}
}

func NewWeiboStatus(id string, text string, url string,  imageurl string, author string) Messager {
    return &WeiboStatus{id, text, url, imageurl}
}

func (w *WeiboStatus) Tokenize(tf TokenFactory, c TokenContainer, mc MessagerContainer) error {
    text := w.Text()
    //tokenize:
    //0. use regex to replace http://
    //1. split the text to sections, using the edge of different rune size, and symbols.
    //2. every connected n runes became a token.

    reg, _ := regexp.Compile(`http:\/\/[a-zA-Z\.\/0-9]*`)
    text = reg.ReplaceAllString(text, "")

    runes := make([]rune, 0, len(text)/5)
    sections := make([]string, 0, len(text)/5)
    state := 0 //0: inital 1:english 2:utf8 runes
    totaltoken := 0
    //step1
    for _, r := range text {
        switch state {
        case 0: //inital
            switch {
            case (r > 'a' && r < 'z' ) || (r > 'A' && r < 'Z'):
                runes = append(runes, r)
                state = 1
            case (r > 0x3400 && r < 0x9fff) || (r > 0x20000 && r < 0x2A6DF):
                runes = append(runes, r)
                state = 2
            default:
                state = 0
            }

        case 1: //english
            switch {
            case (r > 'a' && r < 'z' ) || (r > 'A' && r < 'Z'):
                state = 1
                runes = append(runes, r)
            case (r > 0x3400 && r < 0x9fff) || (r > 0x20000 && r < 0x2A6DF):
                state = 2
                if len(string(runes)) > 3 {
                    c.Active(tf.NewTokener(string(runes),50,w))
                    totaltoken ++
                }
                runes = make([]rune, 0, len(text)/5)
                runes = append(runes, r)
            default:
                state = 0
                //sections = append(sections, string(runes))
                runes = make([]rune, 0, len(text)/5)
            }
        case 2: //utf8
            switch {
            case (r > 'a' && r < 'z' ) || (r > 'A' && r < 'Z'):
                state = 1
                sections = append(sections, string(runes))
                runes = make([]rune, 0, len(text)/5)
                runes = append(runes, r)
            case (r > 0x3400 && r < 0x9fff) || (r > 0x20000 && r < 0x2A6DF):
                state = 2
                runes = append(runes, r)
            default:
                state = 0
                sections = append(sections, string(runes))
                runes = make([]rune, 0, len(text)/5)
            }
        }
    }
    //step2
    for _, section := range sections {
        //fmt.Printf("section: %s", section)
        runes := []rune(section)
        //runelist := list.New()
        runearray := make([]int, len(section))
        for i, r := range runes {
            //fmt.Printf("---%d, %d, %s, %d\n", i, int(r), string(r), len(runearray))
            runearray[i] = int(r)
            //fmt.Println(runearray)
            var postfix string
            if i > 0 {
                postfix = string(rune(runearray[i-1])) + string(r)
                //fmt.Printf("token: %s", postfix)
                c.Active(tf.NewTokener(postfix,7,w))
                totaltoken ++
            }
            if i > 1 {
                postfix = string(rune(runearray[i-2])) + postfix
                c.Active(tf.NewTokener(postfix,11,w))
                totaltoken ++
            }
            if i > 2 {
                postfix = string(rune(runearray[i-3])) + postfix
                c.Active(tf.NewTokener(postfix,13,w))
                totaltoken ++
            }
            if i > 3 {
                postfix = string(rune(runearray[i-4])) + postfix
                c.Active(tf.NewTokener(postfix,17,w))
                totaltoken ++
            }
            if i > 4 {
                postfix = string(rune(runearray[i-5])) + postfix
                c.Active(tf.NewTokener(postfix,19,w))
                totaltoken ++
            }
            if i > 5 {
                postfix = string(rune(runearray[i-6])) + postfix
                c.Active(tf.NewTokener(postfix,23,w))
                totaltoken ++
            }
            if i > 6 {
                postfix = string(rune(runearray[i-7])) + postfix
                c.Active(tf.NewTokener(postfix,29,w))
                totaltoken ++
            }
        }
    }
    if nil != mc {
        mc.Active(w,totaltoken)
    }
    return nil
}

func (w *WeiboStatus) Author() string {
    return ""
}

func (w *WeiboStatus) ImageUrl() string {
    return w.imageUrl
}

func (w *WeiboStatus) Url () string {
    return w.url
}

func (w *WeiboStatus) Text() string {
    return w.text
}

func (w *WeiboStatus) Id() string {
    //fmt.Println(w.idstr)
    return w.idstr
}



