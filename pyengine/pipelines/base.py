import threading
import time

import logging


class BasePipeline(threading.Thread):
    interval = 60

    def __init__(self, *args, **kwargs):
        threading.Thread.__init__(self, name=kwargs.get('name'))
        self.name = kwargs.get('name')

    def run(self):
        logging.info("启动 {} 线程".format(self.name))
        while True:
            try:
                logging.info("执行 {} 的任务".format(self.name))
                self.task()
            except Exception as e:
                logging.error("{} 执行任务发生错误: {}".format(self.name, e))
            time.sleep(self.interval)
