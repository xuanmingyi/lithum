from models.base import Base
from sqlalchemy import Column, String, Integer, DateTime
from datetime import datetime


class LVV2Image(Base):
    __tablename__ = "lvv2_image"

    id = Column(Integer, primary_key=True)
    thread_id = Column(Integer)
    url = Column(String(256))
    status = Column(String(32))
    date = Column(DateTime, default=datetime.now())