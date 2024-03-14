package network

import "testing"

func TestHttp(t *testing.T) {
	request := NewGETRequest("http://www.zhattatey.top", false, nil)
	response, _ := request.Send()
	println(response.Body)
}
