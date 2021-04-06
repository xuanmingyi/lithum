from pipelines.base import BasePipeline

from utils import download_html
from bs4 import BeautifulSoup

from models import LVV2Thread


class LVV2ThreadPipeline(BasePipeline):
    interval = 60*60

    def __init__(self):
        super(LVV2ThreadPipeline, self).__init__(**{"name": "lvv2_thread"})

    def task(self):
        html = download_html("https://lvv2.com/nsfw", proxy=True)
        soup = BeautifulSoup(html, features="html.parser")
        for _thread in soup.find_all("div", class_="link show"):
            url = _thread.find("a", class_="thumbnail")["href"]
            title = _thread.find("a", class_="title").text
            tag = _thread.find("h4").text
            if self.session.query(LVV2Thread).filter(LVV2Thread.url==url).count() == 0:
                self.session.add(LVV2Thread(url=url, status="new", tag=tag, title=title))
                self.session.commit()
