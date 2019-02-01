package net

import (
	"fmt"
	"net/url"
	"testing"
)

func TestHTTPGet(t *testing.T) {
	res, err := HTTPGet("http://www.01happy.com/demo/accept.php?id=1")
	if err == nil {
		println(res)
	}
}

func TestHTTPPost(t *testing.T) {
	res, err := HTTPPost("http://www.01happy.com/demo/accept.php","name=c&age=30")
	if err == nil {
		fmt.Println(res)
	}
}

//HTTPPostForm is function to perform http post form.
func TestHTTPPostForm(t *testing.T) {
	res, err := HTTPPostForm("http://www.01happy.com/demo/accept.php", url.Values{"key": {"Value"}, "id": {"123"}})
	if err == nil {
		fmt.Println(res)
	}
}
