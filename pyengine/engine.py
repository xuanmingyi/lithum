from pipelines import PIPELINES


class Runner(object):
    pipelines = []

    def __init__(self):
        for _class in PIPELINES:
            self.pipelines.append(_class())

    def start(self):
        for pipeline in self.pipelines:
            pipeline.start()

    def join(self):
        for pipeline in self.pipelines:
            pipeline.join()


def main():
    runner = Runner()
    runner.start()
    runner.join()


if __name__ == '__main__':
    main()
