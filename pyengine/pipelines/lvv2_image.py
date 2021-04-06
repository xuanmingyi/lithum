from pipelines.base import BasePipeline

from models import LVV2Thread
from models import LVV2Image

from utils import download_html
from bs4 import BeautifulSoup


class LVV2ImagePipeline(BasePipeline):
    interval = 20 * 60

    def __init__(self):
        super(LVV2ImagePipeline, self).__init__(**{"name": "lvv2_image"})

    def task(self):
        threads = self.session.query(LVV2Thread).filter(LVV2Thread.status == "new").all()
        for _thread in threads:
            html = download_html(_thread.url, proxy=True)
            soup = BeautifulSoup(html, features="html.parser")
            for _image in soup.find_all("img", "lazy detailImg"):
                url = _image["data-echo"]
                if self.session.query(LVV2Image).filter(LVV2Image.url==url, LVV2Image.thread_id==_thread.id).count() == 0:
                    self.session.add(LVV2Image(url=url, status="new", thread_id=_thread.id, date=_thread.date))
                    _thread.status = "download"
                    self.session.commit()
