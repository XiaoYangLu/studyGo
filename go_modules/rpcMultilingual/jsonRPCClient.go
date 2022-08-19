package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct{}

func (h Hello) SayHello(req string, res *string) error {
	*res = req
	return nil
}
func main() {
	//1、注册远程服务
	conn, err := net.Dial("tcp", "192.168.31.205:8080")
	if err != nil {
		fmt.Println(err)
	}
	//2、延迟关闭
	defer conn.Close()
	//基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//3、调用远程函数
	var reply string
	err = client.Call("hello.SayHello", "lisi", &reply)
	if err != nil {
		fmt.Println(err)
	}
	//4、获取远程数据
	fmt.Println(reply)
}
