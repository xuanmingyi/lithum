from pipelines.base import BasePipeline

from models import JiandanImage, SessionFactory
from sqlalchemy.orm import scoped_session
from utils import download_image

from config import Config

import os


class JiandanDownloader(BasePipeline):
    interval = 300

    def __init__(self):
        super(JiandanDownloader, self).__init__(**{"name": "jiandan_downloader"})

    def task(self):
        self.session = scoped_session(SessionFactory)
        images = self.session.query(JiandanImage).filter(JiandanImage.status == "new").all()
        count = 0
        for _image in images:
            download_image(_image.url, os.path.join(Config.Get('default.base_output'), "jiandan", _image.date.strftime("%Y-%m-%d")))
            _image.status = "download"
            self.session.commit()
            count += 1
            self.logger.info('下载图片: {}'.format(_image.url))
        self.logger.info('下载图片: {} 张'.format(count))
        self.session.remove()
