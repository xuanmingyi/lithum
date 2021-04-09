import threading
import time

import logging


class BasePipeline(threading.Thread):
    interval = 60

    def __init__(self, *args, **kwargs):
        threading.Thread.__init__(self, name=kwargs.get('name'))
        self.name = kwargs.get('name')

        self.logger = logging.getLogger(self.name)

    def run(self):
        self.logger.info("启动 {} 线程".format(self.name))
        while True:
            try:
                self.logger.info("执行 {} 的任务".format(self.name))
                self.task()
            except Exception as e:
                self.logger.error("{} 执行任务发生错误: {}".format(self.name, e))
            time.sleep(self.interval)
