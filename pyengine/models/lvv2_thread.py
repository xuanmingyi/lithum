from models.base import Base
from sqlalchemy import Column, String, Integer, DateTime
from datetime import datetime


class LVV2Thread(Base):
    __tablename__ = "lvv2_thread"

    id = Column(Integer, primary_key=True)
    title = Column(String(255))
    url = Column(String(255))
    status = Column(String(32))
    date = Column(DateTime, default=datetime.now())
    tag = Column(String(32))