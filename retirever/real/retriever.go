package real

import (
	"io/ioutil"
	"net/http"
	"time"
)

//实现接口
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
	//req := HttpRequest.NewRequest()
	/*resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpRequest(resp, true)

	//resp.Body.Close()
	if err != nil {
		panic(err)
	}*/

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
