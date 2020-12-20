package main

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RPC RPC `yaml:"rpc"`
}

type RPC struct{
	Type string `yaml:"type"`
	NSQ NSQ `yaml:"nsq"`
}


type NSQ struct {
	NslookupdTcpAddress string `yaml:"nslookupd_tcp_address"`
	Channel string `yaml:"channel"`
	Topic string `yaml:"topic"`
}

var Conf Config

func (c *Config) ReadConfig() error{
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &Conf)
	if err != nil {
		return err
	}

	return nil
}
