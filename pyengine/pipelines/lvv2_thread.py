from pipelines.base import BasePipeline

from bs4 import BeautifulSoup
from sqlalchemy.orm import scoped_session
from datetime import datetime

from models import LVV2Thread, SessionFactory
from utils import download_html


class LVV2ThreadPipeline(BasePipeline):
    interval = 60*60

    def __init__(self):
        super(LVV2ThreadPipeline, self).__init__(**{"name": "lvv2_thread"})

    def task(self):
        self.session = scoped_session(SessionFactory)
        html = download_html("https://lvv2.com/nsfw", proxy=True)
        soup = BeautifulSoup(html, features="html.parser")
        count = 0
        for _thread in soup.find_all("div", class_="link show"):
            url = _thread.find("a", class_="thumbnail")["href"]
            title = _thread.find("a", class_="title").text
            tag = _thread.find("h4").text
            if self.session.query(LVV2Thread).filter(LVV2Thread.url == url).count() == 0:
                self.session.add(LVV2Thread(url=url, status="new", tag=tag, title=title, date=datetime.now()))
                self.session.commit()
                count += 1
        self.logger.info('抓取数据: {} 条'.format(count))
        self.session.remove()
