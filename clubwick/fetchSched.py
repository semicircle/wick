# -*- coding: UTF-8 -*-
import sys
import urllib
import urllib2
import httplib
import HTMLParser
import re
#from BeautifulSoup import BeautifulSoup
from bs4 import BeautifulSoup
from datetime import *
from time import *

class SchedItem:
    title=''
    hostTeam=''
    visitTeam=''
    startTime=''
    endTime=''
    link=''
    pass

def fetch_sched_118(url = 'http://data.188score.com/season/match/5620-2.html'):
    ret = []
    content = urllib2.urlopen(url).read()
    soup = BeautifulSoup(content, from_encoding="GB18030")
    tbl = soup.find(id = 'vs_table')
    trs = tbl.tbody.find_all('tr')
    #print trs 
    for tr in trs:
        print tr
        item = SchedItem()
        item.hostTeam = tr.contents[1].next
        item.visitTeam = tr.contents[3].next
        strtime = tr.contents[0].next
        print item.hostTeam
        print item.visitTeam
        print strtime
        dt = datetime.strptime(strtime, "%Y-%m-%d %H:%M:%S")
        item.startTime = str(int(mktime(dt.timetuple()))) + '000'
        item.endTime = str(int(mktime(dt.timetuple())) + 6000) + '000'
        item.title = dt.strftime("%m-%d ") + item.hostTeam + 'vs' + item.visitTeam
        ret.append()

    return ret



def fetch_sched_zgzcw(url = 'http://www.zgzcw.com/saicheng/yijia.htm'):
    ret = []
    content = urllib2.urlopen(url).read()
    soup = BeautifulSoup(content, from_encoding="GB18030")
    trs = soup.find_all('tr') 
    for tr in trs:
        #if (tr.color != '#f8f8f8'):
        #    continue
        #print '=========tr=========='
        if (not tr.has_key('bgcolor')):
            continue
        if (tr['bgcolor'] == '#F8f8f8'):
            #print 'game line'
            #print tr.contents[0].next
            item = SchedItem()
            item.hostTeam = tr.contents[0].next
            item.visitTeam = tr.contents[1].next
            strtime = tr.contents[2].next
            #print strtime
            dt = datetime.strptime(strtime, "%Y-%m-%d %H:%M:%S")
            item.startTime = str(int(mktime(dt.timetuple()))) + '000'
            item.endTime = str(int(mktime(dt.timetuple())) + 6000) + '000'
            item.title = dt.strftime("%m-%d ") + item.hostTeam + 'vs' + item.visitTeam
            #print item.title
            ret.append(item)

    return ret

fetch_sched_118()
