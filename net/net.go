package net

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//HTTPGet is a function to perform http get.
func HTTPGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("HttpGet error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("HttpGet readAll error")
	}
	return string(body), nil

}

//HTTPPost is a function to perform http post.
//params is key=value format
// account=name&age=3&...
func HTTPPost(requestURL string,params string) (string, error) {
	//要設定成"application/x-www-form-urlencoded"
	resp, err := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(params))
	if err != nil {
		return "", errors.New("HTTPPost error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("HTTPPost readAll error")
	}

	return string(body), nil
}

//HTTPPostForm is function to perform http post form.
func HTTPPostForm(requestURL string, val url.Values) (string, error) {
	resp, err := http.PostForm(requestURL, val)

	if err != nil {
		return "", errors.New("HTTPPostForm error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("HTTPPostForm readAll error")
	}

	return string(body), nil
}


