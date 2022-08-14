package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//step1:确认服务器地址和端口，ip:port
	port := ":9527"
	tcpaddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		fmt.Println(err)
	}
	//step2:获取监听端口
	tcpListener, err := net.ListenTCP("tcp4", tcpaddr)
	if err != nil {
		fmt.Println(err)
	}
	//step3:监听该端口
	fmt.Println("服务端聊天界面。。。。。。。。。。。。。。")
	tcpConn, err := tcpListener.AcceptTCP()
	if err != nil {
		fmt.Println(err)
	}
	//step4:数据交互
	//rTcpConn := make([]byte, 520)
	//lenConn, err := tcpConn.Read(rTcpConn)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("服务端说:", string(rTcpConn[:lenConn]))

	var wg sync.WaitGroup
	wg.Add(2)
	//开启2个go协程，并阻塞其运行，完成多句在线聊天
	go func() {
		//fmt.Println("go11111") //go1、2都处于阻塞状态
		defer wg.Done()
		for {
			//循环写
			wTcpConn := ""
			fmt.Scanln(&wTcpConn) //阻塞
			if wTcpConn == "over" {
				tcpConn.Write([]byte(wTcpConn))
				break
			}
			tcpConn.Write([]byte(wTcpConn))
		}
	}()
	go func() {
		//fmt.Println("go2222")
		defer wg.Done()
		for {
			//循环读
			rTcpConn := make([]byte, 520)
			lenTcpConn, err := tcpConn.Read(rTcpConn)  //阻塞
			if err != nil || string(rTcpConn[:lenTcpConn]) == "over" {
				//fmt.Println("错误1：", err)
				fmt.Println("客户端说:", string(rTcpConn[:lenTcpConn]))
				break
			}
			fmt.Println("客户端说:", string(rTcpConn[:lenTcpConn]))
		}
	}()

	wg.Wait()
	defer tcpConn.Close()
}
