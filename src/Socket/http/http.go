package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//http.Get
	//resp, err := http.Get("https://xueyuanjun.com")
	//if err != nil {
	//	fmt.Printf("发起请求失败：%v", err)
	//	return
	//}

	//http.PostForm
	resp, err := http.PostForm("https://xueyuanjun.com/login", url.Values{"name": {"学院君"}, "password": {"test-passwd"}})

	if err != nil {
		// 处理错误
		return
	}

	if resp.StatusCode != http.StatusOK {
		// 处理错误
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
