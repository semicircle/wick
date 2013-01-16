import sys
import urllib
import httplib
import HTMLParser
import re

class SinaNewsParser(HTMLParser.HTMLParser):
    isInTagA = False
    currTitle = ''
    currLink = ''
    newsList = []
    def __init__(self, pattern):
        HTMLParser.HTMLParser.__init__(self)
        self.prog = re.compile(pattern)
        newsList = []
    def handle_starttag(self, tag, attrs):
        if ('a' == tag):
            #print 'starttag a:'
            #print attrs
            self.isInTagA = True
            for attr in attrs:
                if (attr[0] == 'href'):
                    self.currLink = attr[1] + '&pwt=all'
                    break
                pass
            pass
        pass
    def handle_endtag(self, tag):
        if ('a' == tag):
            #print 'endtag a'
            self.isInTagA = False
            if (None != self.prog.search(self.currLink)):
                self.newsList.append((self.currTitle, self.currLink))
            pass
        pass
    def handle_data(self, data):
        if (self.isInTagA):
            #print data
            self.currTitle = data
        pass
    pass

def fetch_news(club='d5400t13v19'):
    content = urllib.urlopen('http://sports.sina.cn/?sa=' + club + '&vt=4').read()
    #content = unicode(content_src, 'UTF-8')
    parser = SinaNewsParser('t24v4')
    parser.newsList = []
    parser.feed(content)
    parser.close()
    #print '======== newsList result ========='
    return parser.newsList

#fetch_news()
