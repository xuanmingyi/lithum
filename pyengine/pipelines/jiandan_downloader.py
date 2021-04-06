from pipelines.base import BasePipeline

from models import JiandanImage
from utils import download_image

from config import BASE_OUTPUT

import os


class JiandanDownloader(BasePipeline):
    interval = 300

    def __init__(self):
        super(JiandanDownloader, self).__init__(**{"name": "jiandan_downloader"})

    def task(self):
        images = self.session.query(JiandanImage).filter(JiandanImage.status == "new").all()
        for _image in images:
            download_image(_image.url, os.path.join(BASE_OUTPUT, "jiandan", _image.date.strftime("%Y-%m-%d")))
            _image.status = "download"
            self.session.commit()
