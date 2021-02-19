from pipelines.base import BasePipeline

from models.lvv2_image import LVV2Image
from models.lvv2_thread import LVV2Thread
from utils import download_image

import os.path

from config import BASE_OUTPUT

import time


class LVV2ImageDownloaderPipeline(BasePipeline):
    interval = 20 * 60
    threads_cache = {}

    def __init__(self):
        super(LVV2ImageDownloaderPipeline, self).__init__(**{"name": "lvv2_image_dwonloader"})

    def task(self):
        images = self.session.query(LVV2Image).filter(LVV2Image.status=="new").all()
        for _image in images:
            download_image(_image.url, os.path.join(BASE_OUTPUT, "lvv2", str(_image.date), self.get_thread_by_id(_image.thread_id).title))
            #_update_image = self.session.query(LVV2Image).get(_image.id)
            _image.status = "download"
            self.session.commit()
            time.sleep(1)

    def get_thread_by_id(self, _id):
        if str(_id) in self.threads_cache.keys():
            return self.threads_cache[str(_id)]
        obj = self.session.query(LVV2Thread).filter(LVV2Thread.id==_id).first()
        self.threads_cache[str(_id)] = obj
        return self.threads_cache[str(_id)]