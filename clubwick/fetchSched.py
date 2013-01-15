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

#unfort. 118 use javascript to update match table.
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
        item.title = item.hostTeam + 'vs' + item.visitTeam
        ret.append()

    return ret

def fetch_sched_500(url = 'http://liansai.500.com/cupindex-seasonid-2396', 
        baseUrl = 'http://liansai.500.com/'):
    #1. find a select named 'season_list' and use the first option's value
    #   as a new url ..500.com/cupindex-seasonid-XXXX 
    #2. find a iframe named 'fra' and navigate into the src url.
    #3. a table.
    content = urllib2.urlopen(url).read()
    soup = BeautifulSoup(content, from_encoding="GB18030")
    #1.
    select = soup.find(id = 'season_list')
    seasonid = select.contents[0]['value']
    print 'the newest seasonid:' + seasonid
    newUrl = baseUrl + 'cupindex-seasonid-' + seasonid
    if url != newUrl :
        print 'redirect to newer season, url:' + newUrl
        content = urllib2.urlopen(newUrl).read()
        soup = BeautifulSoup(content, from_encoding="GB18030")
    #2.
    iframe = soup.find(id = 'fra')
    fra_src = iframe['src']
    newUrl = baseUrl + fra_src
    print 'redirect to iframe (which contains the match table):'
    print newUrl
    content = urllib2.urlopen(newUrl).read()
    soup = BeautifulSoup(content, from_encoding="GB18030")
    #3.
    trs = [tr for tr in soup.find_all('tr') if 'tr04' in tr['class'] or 'tr03' in tr['class'] ]
    print 'match tr(table row) detected: ' + str(len(trs))
    ret = []
    for tr in trs:
        #print tr.contents[1].string
        #print ' '.join([slice for slice in tr.contents[1].children if slice.string != None])
        item = SchedItem()
        item.hostTeam = tr.contents[3].contents[0].string
        item.visitTeam = tr.contents[7].contents[0].string
        strtime = ' '.join([slice for slice in tr.contents[1].children if slice.string != None])
        #print item.hostTeam
        #print item.visitTeam
        #print strtime
        dt = datetime.strptime(strtime, "%Y-%m-%d %H:%M:%S")
        item.startTime = str(int(mktime(dt.timetuple()))) + '000'
        item.endTime = str(int(mktime(dt.timetuple())) + 6000) + '000'
        item.title = item.hostTeam + 'vs' + item.visitTeam
        ret.append(item)

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
            item.title = item.hostTeam + 'vs' + item.visitTeam
            #print item.title
            ret.append(item)

    return ret

#fetch_sched_118()
#print "TEST 1"
#print fetch_sched_500()
#print "TEST 2"
#print fetch_sched_500(url = 'http://liansai.500.com/cupindex.php?seasonid=1189')
