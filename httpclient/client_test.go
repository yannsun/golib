package httpclient

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T)  {
	url := "http://182.92.159.121:8081/cas/serviceValidate?ticket=ST-1-4bPYjlIXxkkUBZdbRUWO-cas01.example.org&service=https://www.baidu.com/"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(body)
	t.Logf("%s", body)
}
