package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"net"
	"time"
)

func main() {
 i:=9

	for{
		i+=1
	kcpconn, err := kcp.DialWithOptions("localhost:10000", nil, 10, 3)
	if err!=nil {
		panic(err)
	}
	tcpconn,err:=net.Dial("tcp","127.0.0.1:10001")
		if err!=nil {
			fmt.Println(err)
			return
		}

			tcpconn.Write([]byte("hello tcp.emmmmmmmmmmmmmmm"))


			kcpconn.Write([]byte("hello kcp.emmmmmmmmmmmmmmm"))

time.Sleep(2*time.Second)



		fmt.Println(i)

	}
}