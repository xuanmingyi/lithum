DATABASE_DSN="mysql+pymysql://username:password@host:port/dbname"
BASE_OUTPUT="/mnt/all/"

PROXY={
    "http": "http://127.0.0.1:58591",
    "https": "http://127.0.0.1:58591"
}

try:
    from local_config import *
except:
    pass

from pipelines import PIPELINES
