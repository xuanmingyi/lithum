#!/usr/bin/env bash
len=`ps axu|grep main|grep -v grep|wc -l`

if [ $len -eq 0 ];then
	cd /home/pi/gocode/src/github.com/lithum/engine
	nohup ./main >> /home/pi/out.log  2>&1 &
fi
