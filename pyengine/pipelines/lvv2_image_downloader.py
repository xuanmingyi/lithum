from pipelines.base import BasePipeline

from sqlalchemy.orm import scoped_session

from config import Config
import os.path
import time


from models import LVV2Image, LVV2Thread, SessionFactory
from utils import download_image


class LVV2ImageDownloaderPipeline(BasePipeline):
    interval = 60
    threads_cache = {}

    def __init__(self):
        super(LVV2ImageDownloaderPipeline, self).__init__(**{"name": "lvv2_image_downloader"})

    def task(self):
        self.session = scoped_session(SessionFactory)
        images = self.session.query(LVV2Image).filter(LVV2Image.status == "new").all()
        download_count = 0
        notfound_count = 0
        exception_count = 0
        for _image in images:
            try:
                status_code = download_image(_image.url, os.path.join(Config.Get('default.base_output'), "lvv2", _image.date.strftime("%Y-%m-%d"),
                                                                      self.get_thread_by_id(_image.thread_id).title))
                if status_code == 200:
                    _image.status = 'download'
                    self.logger.info('下载图片: {}'.format(_image.url))
                    download_count += 1
                if status_code == 404:
                    _image.status = 'notfound'
                    notfound_count += 1
            except Exception as e:
                _image.status = str(e)[:16]
                exception_count += 1
            self.session.commit()
            time.sleep(1)
        self.logger.info('下载图片: 成功 {} 张, Not Found {} 张, 错误 {} 张'.format(download_count, notfound_count, exception_count))
        self.session.close()

    def get_thread_by_id(self, _id):
        if str(_id) in self.threads_cache.keys():
            return self.threads_cache[str(_id)]
        obj = self.session.query(LVV2Thread).filter(LVV2Thread.id == _id).first()
        self.threads_cache[str(_id)] = obj
        return self.threads_cache[str(_id)]
