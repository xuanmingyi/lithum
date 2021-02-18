import threading
import time

from models.base import Session


class BasePipeline(threading.Thread):
    interval = 60

    def __init__(self, *args, **kwargs):
        threading.Thread.__init__(self, name=kwargs.get('name'))
        self.name = kwargs.get('name')
        self.path = kwargs.get('path')
        self.session = Session()

    def run(self):
        while True:
            self.task()
            time.sleep(self.interval)