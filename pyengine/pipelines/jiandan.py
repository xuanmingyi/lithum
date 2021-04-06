from pipelines.base import BasePipeline

from utils import download_html
from bs4 import BeautifulSoup

from models import JiandanImage


class Jiandan(BasePipeline):
    interval = 300

    def __init__(self):
        super(Jiandan, self).__init__(**{"name": "jiandan"})

    def task(self):
        html = download_html("http://jandan.net/girl")
        soup = BeautifulSoup(html, features="html.parser")
        for _image in soup.find_all("a", class_="view_img_link"):
            url = "http:{}".format(_image["href"])
            if self.session.query(JiandanImage).filter(JiandanImage.url == url).count() == 0:
                self.session.add(JiandanImage(url=url, status="new"))
                self.session.commit()
