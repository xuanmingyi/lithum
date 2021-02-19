import errno
import os
import os.path
import re
import requests

from config import PROXY, WHITE_LIST

import logging



def download_html(url, proxy=False):
    headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0",
               "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
               "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
               "Host": re.search(r"://([\w|\.]+)/", url).group(1), "Referer": url}
    if proxy:
        response = requests.get(url, headers=headers, proxies=PROXY)
    else:
        response = requests.get(url, headers=headers)
    if response.status_code != 200:
        raise Exception()
    return response.text


def download_image(url, path, proxy=False):
    while True:
        host = re.search(r"://([\w|\.]+)/", url).group(1)
        headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0",
                   "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
                   "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
                   "Host": host, "Referer": url}
        response = requests.get(url, headers=headers, allow_redirects=False)
        if response.status_code == 301:
            url = response.headers.get("Location")
            continue
        if response.status_code != 200:
            raise Exception("response.status_code 返回值不是 200 : {} ".foramt(url))
        filename = re.search(r"/(\w+\.\w+)$", url).group(1)
        try:
            os.makedirs(path)
        except OSError as exc:
            if exc.errno == errno.EEXIST and os.path.isdir(path):
                pass
            else:
                raise Exception()
        filepath = os.path.join(path, filename)
        with open(filepath, "wb") as f:
            f.write(response.content)
        return