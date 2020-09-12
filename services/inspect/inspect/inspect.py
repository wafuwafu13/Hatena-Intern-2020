import urllib.robotparser

def inspect(url):
    rp = urllib.robotparser.RobotFileParser()
    rp.set_url(url + "/robots.txt")
    rp.read()
    rrate = rp.request_rate("*")
    return rp.can_fetch("*", url)
