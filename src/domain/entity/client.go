package entity

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Client interface {
	Sender(req *http.Request) (*http.Response, error)
}

type client struct {
	httpclient http.Client
}

func NewClient() Client {
	jar, _ := cookiejar.New(nil)
	c := http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &client{c}
}

func (cl *client) init() {

}

func (cl *client) Sender(req *http.Request) (*http.Response, error) {
	resp, err := cl.httpclient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
