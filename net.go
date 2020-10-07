package xmlydownloader

import (
	"net/http"
)

const (
	UserAgentPC      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36"
	UserAgentAndroid = "ting_6.3.60(sdk,Android16)"
)

const (
	PC = iota
	Android
)

var client http.Client

//HttpGet 请求GET
//
//userAgent: PC, Android
func HttpGet(url string, userAgent int) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	switch userAgent {
	case PC:
		req.Header.Set("User-Agent", UserAgentPC)
	case Android:
	default:
		req.Header.Set("User-Agent", UserAgentAndroid)
	}

	return client.Do(req)
}

//HttpGetByCookie 使用Cookie请求GET
//
//userAgent: PC, Android
func HttpGetByCookie(url, cookie string, userAgent int) (*http.Response, error) {
	return get(url, cookie, userAgent)
}

func get(url, cookie string, userAgent int) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	switch userAgent {
	case PC:
		req.Header.Set("User-Agent", UserAgentPC)
	case Android:
	default:
		req.Header.Set("User-Agent", UserAgentAndroid)
	}
	return client.Do(req)
}
