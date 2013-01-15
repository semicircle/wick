import urllib
import sys
import fetchNews
import httplib
import time

#target_server = "semicircle-test.appspot.com"
target_server = "127.0.0.1:1984"

def go():
    newslist = fetchNews.fetch_news()
    conn = httplib.HTTPConnection(target_server)
    #conn = httplib.HTTPConnection("localhost:8080")
    conn.request("GET", "http://semicircle-test.appspot.com/CleanNews")
    #time.sleep(5)
    rsp = conn.getresponse()
    if (200 != rsp.status):
        print rsp.status, rsp.reason
        return
    time.sleep(5)
    print "Start AddNews"
    for item in newslist:
        encpyLink = urllib.quote(item[1])
        conn = httplib.HTTPConnection(target_server)
        params = urllib.urlencode({'title':item[0], 'link':encpyLink})
        time.sleep(2)
        conn.request("POST", "http://semicircle-test.appspot.com/AddNews", params)
        time.sleep(2)
        rsp = conn.getresponse()
        time.sleep(2)
        if (200 != rsp.status):
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
go()





    
