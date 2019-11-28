package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

//http client 应用

func main() {

	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)

	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	//resp, err := http.DefaultClient.Do(request)

	//resp, err := http.Get("http://www.baidu.com")
	client := http.Client{
		//可以设置代理服务器
		//Transport:     nil,
		//重定向会从这个函数里过
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
		//Cookie
		//Jar:     nil,
		Timeout: time.Millisecond * 1000,
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	/*bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	*/

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(bytes))

}
