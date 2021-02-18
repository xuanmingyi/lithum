from models.base import Base
from sqlalchemy import Column, String, Integer, DateTime


class JiandanOOXXImage(Base):
    __tablename__ = "jiandan_ooxx"

    id = Column(Integer, primary_key=True)
    url = Column(String(256))
    status = Column(String(32))
    date = Column(DateTime)
