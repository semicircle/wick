# -*- coding:UTF-8 -*-
import urllib
import sys
import fetchSched
import httplib
import time

#target_server = "127.0.0.1:3000"
target_server = "118.145.12.4:3000"

def update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_acm', pattern=u'AC米兰'):
    schedlist = fetchSched.fetch_sched_zgzcw(url)
    conn = httplib.HTTPConnection(target_server)
    conn.request("GET", "/cleansched/" + club + ".json")
    time.sleep(2)
    rsp = conn.getresponse()
    if (204 != rsp.status):
        print rsp.status, rsp.reason
        return
    time.sleep(2)
    for item in schedlist:
        if ((item.hostTeam != pattern) and (item.visitTeam != pattern)):
            continue
        conn = httplib.HTTPConnection(target_server)
        #print item.startTime
        #print type(club)
        params = urllib.urlencode(
            {'sched_record[title]':item.title.encode('utf8'), 
             'sched_record[link]':'',
             'sched_record[hostteam]':item.hostTeam.encode('utf8'),
             'sched_record[visitteam]':item.visitTeam.encode('utf8'),
             'sched_record[starttime]':item.startTime.encode('utf8'),
             'sched_record[endtime]':item.endTime.encode('utf8'),
             'sched_record[club]':club})

        conn.request("POST", "/sched_records.json", params)
        time.sleep(1)
        rsp = conn.getresponse()
        #time.sleep(5)
        if (201 != rsp.status):
            print rsp.status, rsp.reason
            break
        else:
            print "updating: " + item.title
        pass
    pass

update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_manu', pattern=u'曼彻斯特联')
update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_liv', pattern=u'利物浦')
update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_acm', pattern=u'AC米兰')
update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_int', pattern=u'国际米兰')
update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_juv', pattern=u'尤文图斯')
update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_che', pattern=u'切尔西')
update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_ars', pattern=u'阿森纳')
update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_rma', pattern=u'皇家马德里')
update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_bar', pattern=u'巴塞罗那')

