package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//step1:获取服务器ip+port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "192.168.31.205:9527")
	if err != nil {
		fmt.Println(err)
	}
	//step2:发送连接请求
	tcpConn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
	}
	//step3:数据交互
	//tcpConn.Write([]byte("hello,我是客户端"))

	fmt.Println("客户端聊天界面。。。。。。。。。。。。。。")
	var wg sync.WaitGroup
	wg.Add(2)
	//开启2个go协程，并阻塞其运行，完成多句在线聊天
	go func() {
		//fmt.Println("go11111")
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
		//fmt.Println("go22222")
		defer wg.Done()
		for {
			//循环读
			rTcpConn := make([]byte, 520)
			lenTcpConn, err := tcpConn.Read(rTcpConn) //阻塞
			if err != nil || string(rTcpConn[:lenTcpConn]) == "over" {
				//fmt.Println("错误1：", err)
				fmt.Println("服务端说:", string(rTcpConn[:lenTcpConn]))
				break
			}
			fmt.Println("服务端说:", string(rTcpConn[:lenTcpConn]))
		}
	}()

	//stop4:结束
	wg.Wait()
	defer tcpConn.Close()

}
