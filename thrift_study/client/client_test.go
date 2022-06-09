package client__test

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"go_learning/thrift_study/idl/gen-go/Sample"
	"testing"
)


var ctx = context.Background()

func GetClient() *Sample.GreeterClient {
	addr := ":9090"
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := Sample.NewGreeterClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

//GetUser
func TestGetUser(t *testing.T) {
	client := GetClient()
	rep, err := client.GetUser(ctx, 100)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}

//SayHello
func TestSayHello(t *testing.T) {
	client := GetClient()

	var address string = "address"
	user := &Sample.User{}
	user.Name = "thrift"
	user.Address = &address

	rep, err := client.SayHello(ctx, user)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}

