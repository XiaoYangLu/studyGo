package requesthtml

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
 * 模仿浏览器发起请求,返回html.body
 */
func RequestHtml(url string) string {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("请求错误：", err.Error())
	}
	//return client,request
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//request.Header.Add("Accept-Encoding","gzip, deflate, br")
	//request.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Connection", "keep-alive")
	//request.Header.Add("Cookie","ll="108296"; bid=hRSgcXLrrgU; __utma=30149280.300490593.1659177740.1659177740.1659177740.1; __utmc=30149280; __utmz=30149280.1659177740.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmt=1; ap_v=0,6.0; __utmb=30149280.3.10.1659177740; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1659177826%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_ses.100001.4cf6=*; __utma=223695111.388756746.1659177826.1659177826.1659177826.1; __utmb=223695111.0.10.1659177826; __utmc=223695111; __utmz=223695111.1659177826.1.1.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=D967AC9161B76DF25BBAAF4C02097C6B2|e0ee25e01cd282c27327e710cb09551b; _pk_id.100001.4cf6=74c13910b3336ccf.1659177826.1.1659177923.1659177826.")
	request.Header.Add("Host", "movie.douban.com")
	//request.Header.Add("sec-ch-ua",".Not/A)Brand";v="99", "Google Chrome";v="103", "Chromium";v="103")
	//request.Header.Add("sec-ch-ua-mobile","?0")
	//request.Header.Add("sec-ch-ua-platform",""Windows"")
	//request.Header.Add("Sec-Fetch-Dest","document")
	//request.Header.Add("Sec-Fetch-Mode","navigate")
	//request.Header.Add("Sec-Fetch-Site","none")
	//request.Header.Add("Sec-Fetch-User","?1")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	responet, err := client.Do(request)
	if err != nil {
		fmt.Println("打开错误：", err)
	}
	htmlbyte, err := ioutil.ReadAll(responet.Body)
	if err != nil {
		fmt.Println("读取错误：", err)
	}
	return string(htmlbyte)
}
