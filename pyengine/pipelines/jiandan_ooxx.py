from pipelines.base import BasePipeline

from utils import download_html
from bs4 import BeautifulSoup

from models.jiandan_ooxx_image import JiandanOOXXImage


class JiandanOOXX(BasePipeline):
    interval = 300

    def task(self):
        html = download_html('http://jandan.net/ooxx')
        soup = BeautifulSoup(html, features="html.parser")
        for _image in soup.find_all("a", class_="view_img_link"):
            url = "http:{}".format(_image["href"])
            if self.session.query(JiandanOOXXImage).filter(JiandanOOXXImage.url==url).count() == 0:
                self.session.add(JiandanOOXXImage(url=url, status="new", date="2021-02-02"))
                self.session.commit()