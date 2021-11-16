package broker

import (
	// "fmt"
	"net"

	"github.com/notarock/pobesob/internal/socket"
	"github.com/notarock/pobesob/internal/topic"
)


type Broker struct {
	topics []topic.Topic
}

func InitBroker(subPort, pubPort string) Broker  {
	return Broker{
		topics: []topic.Topic{},
	}
}

func handlePub(conn net.Conn){

}

func handleSub(conn net.Conn){

}
