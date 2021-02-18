from pipelines.base import BasePipeline

from models.jiandan_ooxx_image import JiandanOOXXImage
from utils import download_image

from config import BASE_OUTPUT

import os

class JiandanOOXXDownloader(BasePipeline):
    interval = 300

    def task(self):
        images = self.session.query(JiandanOOXXImage).filter(JiandanOOXXImage.status=="new").all()
        for _image in images:
            download_image(_image.url, os.path.join(BASE_OUTPUT, "jiandan_ooxx", str(_image.date)))
            _image.status = "download"
            self.session.commit()
