# -*- coding:UTF-8 -*-
import urllib
import sys
import fetchSched
import httplib
import time

#target_server = "127.0.0.1:3000"
target_server = "118.145.12.4:3000"

#shortcuts:
#prefix:
PREFIX_CHAMP_C = u'欧冠:'
PREFIX_EURO_C = u'欧联:'
PREFIX_ENG_L = u'英超:'
PREFIX_ITA_L = u'意甲:'
PREFIX_ITA_C = u'意杯:'
PREFIX_ENG_C = u'足总杯:'
#URL
URL_ZGZCW_ENG_L = 'http://www.zgzcw.com/saicheng/yingchao.htm'
URL_ZGZCW_ITA_L = 'http://www.zgzcw.com/saicheng/yijia.htm'
URL_500_CHAMP_C = 'http://liansai.500.com/cupindex-seasonid-2559'
URL_500_EURO_C = 'http://liansai.500.com/cupindex-seasonid-2551'
URL_500_ITA_C = 'http://liansai.500.com/cupindex-seasonid-2638'
URL_500_ENG_C = 'http://liansai.500.com/cupindex-seasonid-2689'
#club
CLUB_ACM='ce2_club_acm'
CLUB_MANU='ce2_club_manu'
CLUB_LIV='ce2_club_liv'
#pattern
PATTERN_ACM=u'AC米兰'
PATTERN_MANU=u'曼彻斯特联'
PATTERN_LIV=u'利物浦'

class League:
    def __init__(self, prefix, url):
        self.prefix = prefix
        self.url = url
    pass

Leag_Eng_L = League(PREFIX_ENG_L, URL_ZGZCW_ENG_L)
Leag_Eng_C = League(PREFIX_ENG_C, URL_500_ENG_C)
Leag_Ita_L = League(PREFIX_ITA_L, URL_ZGZCW_ITA_L)
Leag_Ita_C = League(PREFIX_ITA_C, URL_500_ITA_C)
Leag_Euro_C = League(PREFIX_EURO_C, URL_500_EURO_C)
Leag_Champ_C = League(PREFIX_CHAMP_C, URL_500_CHAMP_C)
Common_Eng = [Leag_Eng_L, Leag_Eng_C, Leag_Euro_C, Leag_Champ_C]
Common_Ita = [Leag_Ita_L, Leag_Ita_C, Leag_Euro_C, Leag_Champ_C]


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




def update_club_sched(club, pattern, leagueList):
    schedlist = []
    for league in leagueList:
        if '500.com' in league.url:
            fetch = fetchSched.fetch_sched_500
        elif 'zgzcw.com' in league.url:
            fetch = fetchSched.fetch_sched_zgzcw
        else:
            print 'Invalid URL: ' + league.url
        toExtend = fetch(league.url)
        if toExtend != None and len(toExtend) != 0:
            for item in toExtend:
                item.title = league.prefix + item.title
            schedlist.extend(toExtend)
        pass
    #check
    if len(schedlist) == 0:
        print 'nothing in schedlist, the "fetch" step seems not working'
        return

    #sort.
    schedlist.sort(cmp = lambda item1, item2 : int(long(item1.startTime[0:-5]) - long(item2.startTime[0:-5])))
    #for item in schedlist:
    #    print item.startTime + item.title

    #clean_sched
    if not clean_sched(club):
        print 'clean sched failed'
        return

    #add_sched_list
    add_sched_list(prefix='', schedlist=schedlist, club=club, pattern=pattern)

    pass


def do_test_update():
    #update_sched(prefix,  url,  club, pattern)
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_int', pattern=u'国际米兰')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yijia.htm', club='ce2_club_juv', pattern=u'尤文图斯')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_che', pattern=u'切尔西')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/yingchao.htm', club='ce2_club_ars', pattern=u'阿森纳')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_rma', pattern=u'皇家马德里')
    #update_sched(url = 'http://www.zgzcw.com/saicheng/xijia.htm', club='ce2_club_bar', pattern=u'巴塞罗那')
    update_club_sched(CLUB_ACM, PATTERN_ACM, Common_Ita) 
    pass

def do_real_update():

    update_club_sched(CLUB_MANU, PATTERN_MANU, Common_Eng) 
    update_club_sched(CLUB_ACM, PATTERN_ACM, Common_Ita) 
    update_club_sched(CLUB_LIV, PATTERN_LIV, Common_Eng)

    pass

#do_test_update()
do_real_update()
