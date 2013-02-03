package main

import (
    "testing"
    "fmt"
//    "unicode/utf8"
    "regexp"
)

//A normal weibo without repost section.
func TestCreateFromJson1(t *testing.T) {
    jsonText := `{"created_at":"Sun Jan 20 11:26:16 +0800 2013","id":3536506746334010,"mid":"3536506746334010","idstr":"3536506746334010",
        "text":"【维基百科搬迁数据中心：下周三天或短暂中断】维基百科东家维基传媒基金会宣布，1月22日、23日、24日三天的17点到第二天凌晨一天（格里尼治时间），维基百科服务将可能中断。此次，维基百科将搬迁到位于弗吉尼亚州阿什伯（Ashburn）的一座新数据中心。http://t.cn/zjFdWLe","source":"<a href=\"http://app.weibo.com/t/feed/1sxHP2\" rel=\"nofollow\">专业版微博</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww2.sinaimg.cn/thumbnail/61e61e8cjw1e0zvar8n7yj.jpg",
        "bmiddle_pic":"http://ww2.sinaimg.cn/bmiddle/61e61e8cjw1e0zvar8n7yj.jpg","original_pic":"http://ww2.sinaimg.cn/large/61e61e8cjw1e0zvar8n7yj.jpg","geo":null,"user":{"id":1642471052,"idstr":"1642471052","screen_name":"TechWeb","name":"TechWeb","province":"11","city":"8","location":"北京 海淀区","description":"新媒体、新技术、新商业互动交流平台。TechWeb.com.cn","url":"http://www.techweb.com.cn/","profile_image_url":"http://tp1.sinaimg.cn/1642471052/50/22822692178/1","profile_url":"techweb","domain":"techweb","weihao":"","gender":"m","followers_count":290897,"friends_count":1499,"statuses_count":15736,"favourites_count":274,"created_at":"Fri Aug 28 09:55:35 +0800 2009","following":true,"allow_all_act_msg":true,"geo_enabled":true,"verified":true,"verified_type":5,"remark":"","allow_all_comment":true,"avatar_large":"http://tp1.sinaimg.cn/1642471052/180/22822692178/1","verified_reason":"TechWeb官方微博","follow_me":false,"online_status":0,"bi_followers_count":1015,"lang":"zh-cn","star":0,"mbtype":0,"mbrank":0,"block_word":0},"reposts_count":15,"comments_count":2,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)
    if msg.Text() != "【维基百科搬迁数据中心：下周三天或短暂中断】维基百科东家维基传媒基金会宣布，1月22日、23日、24日三天的17点到第二天凌晨一天（格里尼治时间），维基百科服务将可能中断。此次，维基百科将搬迁到位于弗吉尼亚州阿什伯（Ashburn）的一座新数据中心。http://t.cn/zjFdWLe" {
        t.Errorf("msg.text: %s", msg.Text())
    }
    if msg.ImageUrl() != "http://ww2.sinaimg.cn/bmiddle/61e61e8cjw1e0zvar8n7yj.jpg" {
        t.Errorf("msg.imageUrl: %s", msg.ImageUrl())
    }
}

//A weibo with repost section
func TestCreateFromJson2(t *testing.T) {
    jsonText := `{"created_at":"Sun Jan 20 11:26:29 +0800 2013","id":3536506801223177,"mid":"3536506801223177","idstr":"3536506801223177",
        "text":"缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑]","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"retweeted_status":{"created_at":"Sun Jan 20 04:40:41 +0800 2013","id":3536404678606628,"mid":"3536404678606628","idstr":"3536404678606628","text":"手哥，您看这位帅哥能给几分？  @留几手","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww4.sinaimg.cn/thumbnail/5947cfb7jw1e0zjktmfqfj.jpg",
        "bmiddle_pic":"http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg","original_pic":"http://ww4.sinaimg.cn/large/5947cfb7jw1e0zjktmfqfj.jpg","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"reposts_count":134,"comments_count":82,"attitudes_count":3,"mlevel":0,"visible":{"type":0,"list_id":0}},"reposts_count":12,"comments_count":8,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)
    if msg.Text() != "缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑] //@夏商: 手哥，您看这位帅哥能给几分？  @留几手" {
        t.Errorf("msg.text: %s", msg.Text())
    }
    if msg.ImageUrl() != "http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg" {
        t.Errorf("msg.imageUrl: %s", msg.ImageUrl())
    }
}

//Tokenizer.
func TestTokenizer(t *testing.T) {
    var tr Tokenizer
    jsonText := `{"created_at":"Sun Jan 20 11:26:29 +0800 2013","id":3536506801223177,"mid":"3536506801223177","idstr":"3536506801223177",
        "text":"缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑]","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"retweeted_status":{"created_at":"Sun Jan 20 04:40:41 +0800 2013","id":3536404678606628,"mid":"3536404678606628","idstr":"3536404678606628","text":"手哥，您看这位帅哥能给几分？  @留几手","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww4.sinaimg.cn/thumbnail/5947cfb7jw1e0zjktmfqfj.jpg",
        "bmiddle_pic":"http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg","original_pic":"http://ww4.sinaimg.cn/large/5947cfb7jw1e0zjktmfqfj.jpg","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"reposts_count":134,"comments_count":82,"attitudes_count":3,"mlevel":0,"visible":{"type":0,"list_id":0}},"reposts_count":12,"comments_count":8,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)
    tr = msg
    //tr.Tokenize()
    tr = tr
}

//Messager.
func TestMessager(t *testing.T) {
    var m Messager
    jsonText := `{"created_at":"Sun Jan 20 11:26:29 +0800 2013","id":3536506801223177,"mid":"3536506801223177","idstr":"3536506801223177",
        "text":"缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑]","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"retweeted_status":{"created_at":"Sun Jan 20 04:40:41 +0800 2013","id":3536404678606628,"mid":"3536404678606628","idstr":"3536404678606628","text":"手哥，您看这位帅哥能给几分？  @留几手","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww4.sinaimg.cn/thumbnail/5947cfb7jw1e0zjktmfqfj.jpg",
        "bmiddle_pic":"http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg","original_pic":"http://ww4.sinaimg.cn/large/5947cfb7jw1e0zjktmfqfj.jpg","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"reposts_count":134,"comments_count":82,"attitudes_count":3,"mlevel":0,"visible":{"type":0,"list_id":0}},"reposts_count":12,"comments_count":8,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)
    m = msg
    if m.Text() != "缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑] //@夏商: 手哥，您看这位帅哥能给几分？  @留几手" {
        t.Errorf("msg.text: %s", msg.Text())
    }
    if m.ImageUrl() != "http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg" {
        t.Errorf("msg.imageUrl: %s", msg.ImageUrl())
    }
}

func TestTokenizeFuncionality(t *testing.T) {
    var tr Tokenizer
    jsonText := `{"created_at":"Sun Jan 20 11:26:29 +0800 2013","id":3536506801223177,"mid":"3536506801223177","idstr":"3536506801223177",
        "text":"缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑]","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"retweeted_status":{"created_at":"Sun Jan 20 04:40:41 +0800 2013","id":3536404678606628,"mid":"3536404678606628","idstr":"3536404678606628","text":"手哥，您看这位帅哥能给几分？  @留几手","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww4.sinaimg.cn/thumbnail/5947cfb7jw1e0zjktmfqfj.jpg",
        "bmiddle_pic":"http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg","original_pic":"http://ww4.sinaimg.cn/large/5947cfb7jw1e0zjktmfqfj.jpg","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"reposts_count":134,"comments_count":82,"attitudes_count":3,"mlevel":0,"visible":{"type":0,"list_id":0}},"reposts_count":12,"comments_count":8,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)
    tr = msg
    factory := getFactory("simple")
    container := factory.NewTokenContainer()
    tr.Tokenize(factory, container, nil )
    if container.CountAll() == 0 {
        t.Errorf("TokenContainer has 0 element after tokenize a weibo text")
    }
    all := container.All()
    //check points
    cp1, cp2, cp3 := false, false, false

    for _, tok := range all {
        //fmt.Println(tok.Text())
        if tok.Text() == "滚粗" {
            cp1 = true
        }
        if tok.Text() == "藏的都是刀" {
            cp2 = true
        }
        if tok.Text() == "眼睛不大心眼不" {
            cp3 = true
        }
    }

    if !(cp1 && cp2 && cp3) {
        t.Errorf("tokenize failed.")
    }
    //t.Log(container.All())
    //fmt.Println(container.All())
    //fmt.Println(container.All()[0].Text())
}

func TestRune(t *testing.T) {
    jsonText := `{"created_at":"Sun Jan 20 11:26:29 +0800 2013","id":3536506801223177,"mid":"3536506801223177","idstr":"3536506801223177",
        "text":"缺了滚粗两字，没得手哥真传。//@土豆苗: 摹仿一下：男人女相，眼睛不大心眼不小，笑得灿烂，背后藏的都是刀。嘴上那颗痦子不详，有被刨坟的征兆。负分！[偷笑]","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"retweeted_status":{"created_at":"Sun Jan 20 04:40:41 +0800 2013","id":3536404678606628,"mid":"3536404678606628","idstr":"3536404678606628","text":"手哥，您看这位帅哥能给几分？  @留几手","source":"<a href=\"http://app.weibo.com/t/feed/3Nrvbn\" rel=\"nofollow\">三星GalaxyNote</a>","favorited":false,"truncated":false,"in_reply_to_status_id":"","in_reply_to_user_id":"","in_reply_to_screen_name":"","thumbnail_pic":"http://ww4.sinaimg.cn/thumbnail/5947cfb7jw1e0zjktmfqfj.jpg",
        "bmiddle_pic":"http://ww4.sinaimg.cn/bmiddle/5947cfb7jw1e0zjktmfqfj.jpg","original_pic":"http://ww4.sinaimg.cn/large/5947cfb7jw1e0zjktmfqfj.jpg","geo":null,"user":{"id":1497878455,"idstr":"1497878455","screen_name":"夏商","name":"夏商","province":"31","city":"6","location":"上海 静安区","description":"小说家。著有四卷本文集《夏商自选集》。个人官微：@夏商读友会 // 普茶客创始人。中国高端普洱茶原创品牌。企业官微：@普茶客官方微博","url":"http://blog.sina.com.cn/xiashang1969","profile_image_url":"http://tp4.sinaimg.cn/1497878455/50/40011372724/1","cover_image":"http://ww1.sinaimg.cn/crop.0.0.980.300/5947cfb7gw1e0wontrm8dj.jpg","profile_url":"xiashang1969","domain":"xiashang1969","weihao":"","gender":"m","followers_count":109953,"friends_count":209,"statuses_count":854,"favourites_count":25,"created_at":"Sat Apr 03 14:28:32 +0800 2010","following":true,"allow_all_act_msg":false,"geo_enabled":false,"verified":true,"verified_type":0,"remark":"","allow_all_comment":true,"avatar_large":"http://tp4.sinaimg.cn/1497878455/180/40011372724/1","verified_reason":"小说家，代表作有《东岸纪事》《乞儿流浪记》","follow_me":false,"online_status":0,"bi_followers_count":208,"lang":"zh-tw","star":0,"mbtype":12,"mbrank":1,"block_word":0},"reposts_count":134,"comments_count":82,"attitudes_count":3,"mlevel":0,"visible":{"type":0,"list_id":0}},"reposts_count":12,"comments_count":8,"attitudes_count":0,"mlevel":0,"visible":{"type":0,"list_id":0}}
        `
    msg := NewWeiboStatusFromJson(jsonText)

    //fmt.Println(msg.Text())
    //r, size := utf8.DecodeRuneInString("")
    //fmt.Println(size)
    //fmt.Println(r)
    //fmt.Println(utf8.RuneCountInString(msg.Text()))

    n := 0
    text := msg.Text()
    runes := make ([]rune, len(text))
    for _,r := range text{
        runes[n] = r
        n++
        //fmt.Printf("%s%d=%x ", string(r),i,r)
    }
    runes = runes[0:n]
    //fmt.Println(string(runes))
}

func TestRegex(t *testing.T) {
    text := `幼升小强化班如同现代“病梅馆” 】教育有病，合该孩子吃药?当热名小学——知名中学——著名大学——好工作成为一道道密切相关的“门槛” 时，老师与家长都唯恐自己的孩子“输”在潮流之后，整个社会也就不可避免地处于一种教育的“病态”。 http://t.cn/zYyDMKH #第一时评#`

    reg, err := regexp.Compile(`http:\/\/[a-zA-Z\.\/0-9]* `)
    if nil != err {
        fmt.Println(err)
    }

    text = reg.ReplaceAllString(text, "")

    //fmt.Println(text)

}
