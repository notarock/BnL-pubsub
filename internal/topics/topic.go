package topic

import (
	"net"
)

type Topic struct {
	name string
	publishers []net.Conn
	subscribers []net.Conn
}

func AddPublisher(tp Topic, conn net.Conn) {
	tp.publishers = append(tp.publishers, conn)
}

func AddSubscriber(tp Topic, conn net.Conn) {
	tp.subscribers = append(tp.subscribers, conn)
}

