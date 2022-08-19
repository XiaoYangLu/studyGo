package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct{}

func (h Hello) SayHello(req string, res *string) error {
	*res = "hello" + req
	return nil
}
func main() {
	//1、注册rpc服务
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	//2、监听端口
	listen, err := net.Listen("tcp", "192.168.31.205:8080")
	if err != nil {
		fmt.Println(err)
	}
	//3、延迟关闭
	defer listen.Close()
	for {
		//4、建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//5、绑定服务
		//rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
