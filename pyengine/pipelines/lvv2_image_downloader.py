from pipelines.base import BasePipeline

from models import LVV2Image
from models import LVV2Thread
from utils import download_image

import os.path

from config import BASE_OUTPUT

import time


class LVV2ImageDownloaderPipeline(BasePipeline):
    interval = 60
    threads_cache = {}

    def __init__(self):
        super(LVV2ImageDownloaderPipeline, self).__init__(**{"name": "lvv2_image_dwonloader"})

    def task(self):
        images = self.session.query(LVV2Image).filter(LVV2Image.status == "new").all()
        for _image in images:
            try:
                status_code = download_image(_image.url, os.path.join(BASE_OUTPUT, "lvv2", _image.date.strftime("%Y-%m-%d"),
                                                                      self.get_thread_by_id(_image.thread_id).title))
                if status_code == 200:
                    _image.status = "download"
                if status_code == 404:
                    _image.status = "notfound"
            except Exception as e:
                _image.status = str(e)[:16]
            self.session.commit()
            time.sleep(1)

    def get_thread_by_id(self, _id):
        if str(_id) in self.threads_cache.keys():
            return self.threads_cache[str(_id)]
        obj = self.session.query(LVV2Thread).filter(LVV2Thread.id == _id).first()
        self.threads_cache[str(_id)] = obj
        return self.threads_cache[str(_id)]