# -*- coding:UTF-8 -*-
import urllib
import sys
import fetchSched
import httplib
import time

#target_server = "127.0.0.1:3000"
target_server = "118.145.12.4:3000"

def clean_sched(club):
    conn = httplib.HTTPConnection(target_server)
    conn.request("GET", "/cleansched/" + club + ".json")
    time.sleep(2)
    rsp = conn.getresponse()
    if (204 != rsp.status):
        print rsp.status, rsp.reason
        return False
    time.sleep(2)
    return True

def add_sched_list(prefix, schedlist, club, pattern):
    for item in schedlist:
        if ((item.hostTeam != pattern) and (item.visitTeam != pattern)):
            continue
        conn = httplib.HTTPConnection(target_server)
        #print item.startTime
        #print type(club)
        params = urllib.urlencode(
            {'sched_record[title]': (prefix + item.title).encode('utf8'), 
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


def update_sched(prefix = u'', url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_acm', pattern=u'AC米兰'):
    schedlist = fetchSched.fetch_sched_zgzcw(url)
    #do clean
    if (schedlist != 0):
        if not clean_sched(club):
            return
        pass
    add_sched_list(prefix, schedlist, club, pattern)

def update_sched2(prefix = u'', url = 'http://liansai.500.com/cupindex-seasonid-2559', club='ce2_club_acm', pattern=u'AC米兰'):
    schedlist = fetchSched.fetch_sched_500(url)
    add_sched_list(prefix, schedlist, club, pattern)

#shortcuts:
PREFIX_CHAMP_L = u'欧冠:'
URL_500_CHAMP_L = 'http://liansai.500.com/cupindex-seasonid-2559'
PREFIX_EURO_L = u'欧联:'
URL_500_EURO_L = 'http://liansai.500.com/cupindex-seasonid-2551'
PREFIX_YINGCHAO = u'英超:'
URL_ZGZCW_YINGCHAO = 'http://www.zgzcw.com/saicheng/yingchao.htm'
PREFIX_YIJIA = 
URL_ZGZCW_YIJIA = 'http://www.zgzcw.com/saicheng/yijia.htm',


def do_test_update():
    update_sched2(url =  club='ce2_club_acm', pattern=u'AC米兰')

def do_real_update():
    #manu,
    update_sched(prefix =  url =  club='ce2_club_manu', pattern=u'曼彻斯特联')
    update_sched2(prefix = , url = 'http://liansai.500.com/cupindex-seasonid-2559', club='ce2_club_manu', pattern=u'曼彻斯特联')
    update_sched2(prefix =  url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_liv', pattern=u'利物浦')

    update_sched(prefix = u'英超:', url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_liv', pattern=u'利物浦')
    update_sched2(prefix = u'欧冠:', url = 'http://liansai.500.com/cupindex-seasonid-2559', club='ce2_club_liv', pattern=u'利物浦')
    update_sched2(prefix = u'欧联:', url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_liv', pattern=u'利物浦')
    update_sched(url =  club='ce2_club_acm', pattern=u'AC米兰')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_int', pattern=u'国际米兰')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_juv', pattern=u'尤文图斯')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_che', pattern=u'切尔西')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_ars', pattern=u'阿森纳')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_rma', pattern=u'皇家马德里')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_bar', pattern=u'巴塞罗那')

#do_test_update()
do_real_update()
