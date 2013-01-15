import urllib
import sys
import fetchNews
import httplib
import time

#target_server = "semicircle-test.appspot.com"
#target_server = "127.0.0.1:1984"
#target_server = "127.0.0.1:3000"
target_server = "118.145.12.4:3000"

def update_news(postfix_url='d5400t13v19', club='ce2_club_acm'):
    
    #clean all
    
    newslist = fetchNews.fetch_news(postfix_url)
    conn = httplib.HTTPConnection(target_server)
    conn.request("GET", "http://" + target_server + "/cleanclub/" + club + ".json")
    time.sleep(2)
    rsp = conn.getresponse()
    if (204 != rsp.status): #204 :no_content
        print rsp.status, rsp.reason
        return
    time.sleep(2)
    #return
    print "Start AddNews"
    for item in newslist:
        encpyLink = urllib.quote(item[1])
        conn = httplib.HTTPConnection(target_server)
        params = urllib.urlencode({
            'news_record[title]':item[0], 
            'news_record[link]':encpyLink, 
            'news_record[club]':club})
        #conn.request("POST", "http://semicircle-test.appspot.com/AddNews", params)
        conn.request("POST", "http://" + target_server + "/news_records.json", params)
        time.sleep(1)
        rsp = conn.getresponse()
        #time.sleep(2)
        if (201 != rsp.status): #201 :created
            print rsp.status, rsp.reason
            break
        else:
            print "updating: " + item[0]
        pass
    pass

#for i in range(0, 1000000):
#    print("Loop for " + str(i) + "times")
#    go()
#    time.sleep(1200)
#    pass

update_news('d5413t13v19', 'ce2_club_manu')
update_news('d5400t13v19', 'ce2_club_acm')
update_news('d5399t13v19', 'ce2_club_int')
update_news('d5411t13v19', 'ce2_club_juv')
#update_news('d5417t13v19', 'ce2_club_manc')
update_news('d5415t13v19', 'ce2_club_che')
update_news('d5414t13v19', 'ce2_club_ars')
update_news('d5416t13v19', 'ce2_club_liv')
update_news('d11651t13v19', 'ce2_club_rma')
update_news('d11652t13v19', 'ce2_club_bar')







    
