from pipelines.jiandan import Jiandan
from pipelines.jiandan_downloader import JiandanDownloader
from pipelines.lvv2_thread import LVV2ThreadPipeline
from pipelines.lvv2_image import LVV2ImagePipeline
from pipelines.lvv2_image_downloader import LVV2ImageDownloaderPipeline

PIPELINES = [Jiandan, JiandanDownloader, LVV2ThreadPipeline, LVV2ImagePipeline, LVV2ImageDownloaderPipeline]
