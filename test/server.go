package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/nsqio/go-nsq"
	"github.com/xuanmingyi/test/protocols"
	"log"
	"os"
	"os/signal"
	"syscall"
)


var Handlers map[string]nsq.HandlerFunc

type ServerHandler struct{}

func (*ServerHandler) DefaultHandle(msg *nsq.Message) (err error) {
	return err
}

func (*ServerHandler) HandleMessage(msg *nsq.Message) (err error) {
	return err
}


func processMessage(body []byte) error {
	var req protocols.AddRequest
	var rsp protocols.AddResponse
	err := proto.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	rsp.Sum = req.Col1 + req.Col2

	client, _ := nsq.NewProducer(Conf.RPC.NSQ.NslookupdTcpAddress, nsq.NewConfig())
	data, err := proto.Marshal(&rsp)
	if err != nil {
		panic(err)
	}
	client.Publish( "sss", data)
	return nil
}

type MyHandler struct {}

func AddHandle(msg *protocols.RPCMessage) (err error) {
	var req protocols.AddRequest
	var rsp *any.Any
	var response protocols.RPCMessage
	var client *nsq.Producer
	var b []byte


	err = ptypes.UnmarshalAny(msg.Details, &req)
	if err != nil {
		panic(err)
	}

	client, err = nsq.NewProducer(Conf.RPC.NSQ.NslookupdTcpAddress, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	rsp, err = ptypes.MarshalAny(&protocols.AddResponse{
		Sum: req.Col1 + req.Col2,
	})
	if  err != nil {
		panic(err)
	}

	response.Details = rsp
	response.MessageId = msg.MessageId
	response.Name = fmt.Sprintf("reply_%s", msg.Name)
	b, err = proto.Marshal(&response)
	if err != nil {
		panic(err)
	}
	client.Publish(fmt.Sprintf("reply_%s#ephemeral", msg.MessageId), b)

	return nil
}


func DefaultHandle(msg *protocols.RPCMessage) (err error) {
	log.Println("默认的操作")
	return nil
}

func (*MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	if len(msg.Body) == 0 {
		return nil
	}

	var rpcMessage protocols.RPCMessage

	if err = proto.Unmarshal(msg.Body, &rpcMessage); err != nil {
		panic(err)
	}

	switch rpcMessage.Name {
	case "add":
		err = AddHandle(&rpcMessage)
		break
	default:
		err = DefaultHandle(&rpcMessage)
		break
	}
	return err
}

type Server struct {
	Consumer *nsq.Consumer
}

var MyServer Server

func MakeServer() (err error){
	MyServer.Consumer, err = nsq.NewConsumer(Conf.RPC.NSQ.Topic, "default", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	MyServer.Consumer.AddHandler(&MyHandler{})
	return nil
}

func Listen() {
	MyServer.Consumer.ConnectToNSQD(Conf.RPC.NSQ.NslookupdTcpAddress)
}

func main() {
	if err := Conf.ReadConfig(); err != nil {
		log.Fatal(err)
		return
	}

	MakeServer()

	Listen()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan	, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
