package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/nsqio/go-nsq"
	uuid "github.com/satori/go.uuid"
	"github.com/xuanmingyi/test/protocols"
	"log"
)

type Waiter struct {}


var data chan int32 = make(chan int32, 1)


func (*Waiter) HandleMessage(msg *nsq.Message) error {
	var rsp protocols.RPCMessage
	var response protocols.AddResponse


	if err := proto.Unmarshal(msg.Body, &rsp); err != nil {
		panic(err)
	}
	if err := ptypes.UnmarshalAny(rsp.Details, &response); err!= nil {panic(err)}
	data <- response.Sum
	return nil
}


func Add(a, b int32) int32 {
	var req protocols.AddRequest
	var steam []byte


	msg_id := uuid.NewV4().String()

	client, err := nsq.NewProducer(Conf.RPC.NSQ.NslookupdTcpAddress, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	req.Col1 = a
	req.Col2 = b

	detail, err := ptypes.MarshalAny(&req)
	if err != nil {panic(err)}

	steam, err = proto.Marshal(&protocols.RPCMessage{
		Name: "add",
		MessageId: msg_id,
		Details: detail,
	})
	if err != nil {
		panic(err)
	}

	client.Publish(Conf.RPC.NSQ.Topic, steam)

	waiter, err := nsq.NewConsumer(fmt.Sprintf("reply_%s#ephemeral", msg_id), "default#ephemeral", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	waiter.AddHandler(&Waiter{})
	waiter.ConnectToNSQD(Conf.RPC.NSQ.NslookupdTcpAddress)
	defer waiter.Stop()

	n := <-data
	return n
}

func main() {
	if err := Conf.ReadConfig(); err != nil {
		log.Fatal(err)
		return
	}

	c := Add(120, 879)

	fmt.Println(c)
}
