package test

import (
	"encoding/json"
	"testing"

	"github.com/zhojiew/appmesh-load/utils"
)

// Test pretty pring json header
func TestPrettyJSON(t *testing.T) {
	data := `{"args": {},"headers": {"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8","Accept-Encoding": "gzip, deflate","Accept-Language": "zh-CN,zh;q=0.9","Connection": "close","Host": "httpbin.org","Upgrade-Insecure-Requests": "1","User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"},"origin": "103.*.*.*","url": "http://httpbin.org/get"}`
	headersJson, _ := json.Marshal(data)
	utils.PrettyJSON(headersJson)
}
