DATABASE_DSN="mysql+pymysql://username:password@host:port/dbname"
BASE_OUTPUT="/mnt/all/"

PROXY={
    "http": "http://127.0.0.1:8118",
    "https": "http://127.0.0.1:8118"
}

PROXY_LIST = []
WHITE_LIST = ["www.xsimg.xyz", "www.smximg.com", "smmimg.com", "www.s2tu.com", "luoimg.com", "sxeimg.com", "skeimg.com",
              "sxotu.com", "www.1024image.com", "www.x6img.com", "images2.imgbox.com", "www.youimg.xyz", "www.vxotu.com",
              "pic606.com", "www.xsspic.com", "www.kanjiantu.com", "skuimg.com", "tuimg.xyz"]


try:
    from local_config import *
except:
    pass
